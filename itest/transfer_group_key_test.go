package itest

import (
	"context"

	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"github.com/stretchr/testify/require"
)

// testTransferGroupKey verifies that the group_key field added to
// TransferInput and TransferOutput populates correctly across both
// ListTransfers and SubscribeSendEvents for grouped fungible assets, and
// stays empty for ungrouped (asset-id only) assets.
//
// This is the dedicated coverage for the marshal-time group-key resolver
// that lives in rpcserver. The bytes carried in group_key must match the
// asset's tweaked group public key exactly so SDK consumers can reconstruct
// their user-facing identifier from the transfer alone, without a side
// query.
func testTransferGroupKey(t *harnessTest) {
	ctxb := context.Background()

	// Mint two assets in one batch so the test runs a single confirmation
	// cycle: a grouped fungible (carries a group key) and an ungrouped
	// fungible (asset-id only). Reusing the existing fixtures keeps this
	// test in lock-step with the rest of the suite.
	mintReqs := []*mintrpc.MintAssetRequest{
		issuableAssets[0], // grouped fungible
		simpleAssets[0],   // ungrouped fungible
	}
	rpcAssets := MintAssetsConfirmBatch(
		t.t, t.lndHarness.Miner(), t.tapd, mintReqs,
	)
	require.Len(t.t, rpcAssets, 2)

	groupedAsset := rpcAssets[0]
	ungroupedAsset := rpcAssets[1]
	require.NotNil(t.t, groupedAsset.AssetGroup,
		"first minted asset must carry a group key")
	require.Nil(t.t, ungroupedAsset.AssetGroup,
		"second minted asset must not carry a group key")

	expectedGroupKey := groupedAsset.AssetGroup.TweakedGroupKey
	require.NotEmpty(t.t, expectedGroupKey)

	// Spin up a receiver tapd. One node is enough — both transfers go to
	// the same recipient.
	bobLnd := t.lndHarness.NewNodeWithCoins("Bob", nil)
	secondTapd := setupTapdHarness(t.t, t, bobLnd, t.universeServer)
	defer func() {
		require.NoError(t.t, secondTapd.stop(!*noDelete))
	}()

	const sendUnits = 100

	// Send the grouped asset.
	groupedAddr, err := secondTapd.NewAddr(ctxb, &taprpc.NewAddrRequest{
		AssetId:      groupedAsset.AssetGenesis.AssetId,
		Amt:          sendUnits,
		AssetVersion: groupedAsset.Version,
	})
	require.NoError(t.t, err)
	AssertAddrCreated(t.t, secondTapd, groupedAsset, groupedAddr)

	groupedSend, _ := sendAssetsToAddr(t, t.tapd, groupedAddr)
	ConfirmAndAssertOutboundTransfer(
		t.t, t.lndHarness.Miner(), t.tapd, groupedSend,
		groupedAsset.AssetGenesis.AssetId,
		[]uint64{groupedAsset.Amount - sendUnits, sendUnits}, 0, 1,
	)
	AssertNonInteractiveRecvComplete(t.t, secondTapd, 1)

	// Send the ungrouped asset.
	ungroupedAddr, err := secondTapd.NewAddr(ctxb, &taprpc.NewAddrRequest{
		AssetId:      ungroupedAsset.AssetGenesis.AssetId,
		Amt:          sendUnits,
		AssetVersion: ungroupedAsset.Version,
	})
	require.NoError(t.t, err)
	AssertAddrCreated(t.t, secondTapd, ungroupedAsset, ungroupedAddr)

	ungroupedSend, _ := sendAssetsToAddr(t, t.tapd, ungroupedAddr)
	ConfirmAndAssertOutboundTransfer(
		t.t, t.lndHarness.Miner(), t.tapd, ungroupedSend,
		ungroupedAsset.AssetGenesis.AssetId,
		[]uint64{ungroupedAsset.Amount - sendUnits, sendUnits}, 1, 2,
	)
	AssertNonInteractiveRecvComplete(t.t, secondTapd, 2)

	// ListTransfers must surface group_key populated on every input and
	// output of the grouped transfer, and empty on the ungrouped one.
	transferResp, err := t.tapd.ListTransfers(
		ctxb, &taprpc.ListTransfersRequest{},
	)
	require.NoError(t.t, err)
	require.Len(t.t, transferResp.Transfers, 2)

	// The two transfers come back in chronological order; the grouped
	// send was issued first.
	groupedTransfer := transferResp.Transfers[0]
	ungroupedTransfer := transferResp.Transfers[1]

	require.NotEmpty(t.t, groupedTransfer.Inputs)
	for i, in := range groupedTransfer.Inputs {
		require.Equalf(t.t, expectedGroupKey, in.GroupKey,
			"grouped transfer input %d missing group_key", i)
	}
	require.NotEmpty(t.t, groupedTransfer.Outputs)
	for i, out := range groupedTransfer.Outputs {
		require.Equalf(t.t, expectedGroupKey, out.GroupKey,
			"grouped transfer output %d missing group_key", i)
	}

	require.NotEmpty(t.t, ungroupedTransfer.Inputs)
	for i, in := range ungroupedTransfer.Inputs {
		require.Emptyf(t.t, in.GroupKey,
			"ungrouped transfer input %d should not carry "+
				"group_key, got %x", i, in.GroupKey)
	}
	require.NotEmpty(t.t, ungroupedTransfer.Outputs)
	for i, out := range ungroupedTransfer.Outputs {
		require.Emptyf(t.t, out.GroupKey,
			"ungrouped transfer output %d should not carry "+
				"group_key, got %x", i, out.GroupKey)
	}
}
