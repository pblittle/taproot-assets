-- name: InsertAddr :one
INSERT INTO addrs (
    version, genesis_asset_id, fam_key, script_key_id, taproot_key_id,
    taproot_output_key, amount, asset_type, creation_time
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id;

-- name: FetchAddrs :many
SELECT 
    version, genesis_asset_id, fam_key, taproot_output_key, amount, asset_type,
    creation_time, managed_from,
    script_keys.tweaked_script_key,
    script_keys.tweak AS script_key_tweak,
    raw_script_keys.raw_key AS raw_script_key,
    raw_script_keys.key_family AS script_key_family,
    raw_script_keys.key_index AS script_key_index,
    taproot_keys.raw_key AS raw_taproot_key, 
    taproot_keys.key_family AS taproot_key_family,
    taproot_keys.key_index AS taproot_key_index
FROM addrs
JOIN script_keys
    ON addrs.script_key_id = script_keys.script_key_id
JOIN internal_keys raw_script_keys
    ON script_keys.internal_key_id = raw_script_keys.key_id
JOIN internal_keys taproot_keys
    ON addrs.taproot_key_id = taproot_keys.key_id
WHERE creation_time >= @created_after
    AND creation_time <= @created_before
    AND (@unmanaged_only = false OR IFNULL(managed_from, true) = @unmanaged_only) 
ORDER BY addrs.creation_time
LIMIT @num_limit OFFSET @num_offset;

-- name: FetchAddrByTaprootOutputKey :one
SELECT
    version, genesis_asset_id, fam_key, taproot_output_key, amount, asset_type,
    creation_time, managed_from,
    script_keys.tweaked_script_key,
    script_keys.tweak AS script_key_tweak,
    raw_script_keys.raw_key as raw_script_key,
    raw_script_keys.key_family AS script_key_family,
    raw_script_keys.key_index AS script_key_index,
    taproot_keys.raw_key AS raw_taproot_key,
    taproot_keys.key_family AS taproot_key_family,
    taproot_keys.key_index AS taproot_key_index
FROM addrs
JOIN script_keys
  ON addrs.script_key_id = script_keys.script_key_id
JOIN internal_keys raw_script_keys
  ON script_keys.internal_key_id = raw_script_keys.key_id
JOIN internal_keys taproot_keys
  ON addrs.taproot_key_id = taproot_keys.key_id
WHERE taproot_output_key = ?;

-- name: SetAddrManaged :exec
WITH target_addr(addr_id) AS (
    SELECT id
    FROM addrs
    WHERE addrs.taproot_output_key = ?
)
UPDATE addrs
SET managed_from = ?
WHERE id = (SELECT addr_id FROM target_addr);

-- name: UpsertAddrEvent :one
WITH target_addr(addr_id) AS (
    SELECT id
    FROM addrs
    WHERE addrs.taproot_output_key = ?
), target_chain_txn(txn_id) AS (
    SELECT txn_id
    FROM chain_txns
    WHERE chain_txns.txid = ?
)
INSERT INTO addr_events (
    creation_time, addr_id, status, chain_txn_id, chain_txn_output_index,
    managed_utxo_id, asset_proof_id, asset_id
) VALUES (
    ?, (SELECT addr_id FROM target_addr), ?,
    (SELECT txn_id FROM target_chain_txn), ?, ?, ?, ?
)
ON CONFLICT (addr_id, chain_txn_id, chain_txn_output_index)
    DO UPDATE SET status = EXCLUDED.status,
                  asset_proof_id = IFNULL(EXCLUDED.asset_proof_id, asset_proof_id),
                  asset_id = IFNULL(EXCLUDED.asset_id, asset_id)
RETURNING id;

-- name: FetchAddrEvent :one
SELECT
    creation_time, status, asset_proof_id, asset_id,
    chain_txns.txid as txid,
    chain_txns.block_height as confirmation_height,
    chain_txn_output_index as output_index,
    managed_utxos.amt_sats as amt_sats,
    managed_utxos.tapscript_sibling as tapscript_sibling,
    internal_keys.raw_key as internal_key
FROM addr_events
LEFT JOIN chain_txns
       ON addr_events.chain_txn_id = chain_txns.txn_id
LEFT JOIN managed_utxos
       ON addr_events.managed_utxo_id = managed_utxos.utxo_id
LEFT JOIN internal_keys
       ON managed_utxos.internal_key_id = internal_keys.key_id
WHERE id = ?;

-- name: QueryEventIDs :many
SELECT
    addr_events.id as event_id, addrs.taproot_output_key as taproot_output_key
FROM addr_events
JOIN addrs
  ON addr_events.addr_id = addrs.id
WHERE addr_events.status >= @status_from 
  AND addr_events.status <= @status_to
  AND IFNULL(@addr_taproot_key, addrs.taproot_output_key) = addrs.taproot_output_key
ORDER by addr_events.creation_time;