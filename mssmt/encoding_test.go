package mssmt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func assertEqualProof(t *testing.T, expected, actual *Proof) {
	t.Helper()

	for i, node := range expected.Nodes {
		other := actual.Nodes[i]
		require.True(t, IsEqualNode(node, other))
	}
}

func assertEqualCompressedProof(t *testing.T, expected, actual *CompressedProof) {
	t.Helper()

	for i, node := range expected.Nodes {
		other := actual.Nodes[i]
		require.True(t, IsEqualNode(node, other))
	}
	require.Equal(t, expected.Bits, actual.Bits)
}

func TestBitPacking(t *testing.T) {
	t.Parallel()

	// Odd number of bits and greater than a byte to test edge case.
	bits := []bool{true, true, false, false, true, false, true, true, false}
	decompressedBits := unpackBits(packBits(bits))

	// Bits up to the expected length should match.
	require.Equal(t, bits, decompressedBits[:len(bits)])

	// Remaining bits should not be set.
	for _, isBitSet := range decompressedBits[len(bits):] {
		require.False(t, isBitSet)
	}
}

func TestProofEncoding(t *testing.T) {
	t.Parallel()

	tree, leaves := randTree(10_000)
	for key := range leaves {
		proof := tree.MerkleProof(key)
		compressed := proof.Compress()

		var buf bytes.Buffer
		err := compressed.Encode(&buf)
		require.NoError(t, err)

		var decodedCompressed CompressedProof
		err = decodedCompressed.Decode(bytes.NewReader(buf.Bytes()))
		require.NoError(t, err)
		assertEqualCompressedProof(t, compressed, &decodedCompressed)

		decodedProof := decodedCompressed.Decompress()
		assertEqualProof(t, proof, decodedProof)
		assertEqualProof(t, proof, decodedProof.Copy())
	}
}
