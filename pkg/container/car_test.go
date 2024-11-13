package container

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCarRoundTrip(t *testing.T) {
	// this car file is a complex and legal CARv1 file
	original, err := os.ReadFile("testdata/sample-v1.car")
	require.NoError(t, err)

	roots, it, err := readCar(bytes.NewReader(original))
	require.NoError(t, err)

	var blks []carBlock
	for blk, err := range it {
		require.NoError(t, err)
		blks = append(blks, blk)
	}

	require.Len(t, blks, 1049)

	buf := bytes.NewBuffer(nil)

	err = writeCar(buf, roots, func(yield func(carBlock, error) bool) {
		for _, blk := range blks {
			if !yield(blk, nil) {
				return
			}
		}
	})
	require.NoError(t, err)

	// Bytes equal after the round-trip
	require.Equal(t, original, buf.Bytes())
}

func FuzzCarRoundTrip(f *testing.F) {
	// Note: this fuzzing is somewhat broken.
	// After some time, the fuzzer discover that a varint can be serialized in different
	// ways that lead to the same integer value. This means that the CAR format can have
	// multiple legal binary representation for the exact same data, which is what we are
	// trying to detect here. Ideally, the format would be stricter, but that's how things
	// are.

	example, err := os.ReadFile("testdata/sample-v1.car")
	require.NoError(f, err)

	f.Add(example)

	f.Fuzz(func(t *testing.T, data []byte) {
		roots, blocksIter, err := readCar(bytes.NewReader(data))
		if err != nil {
			// skip invalid binary
			t.Skip()
		}

		// reading all the blocks, which force reading and verifying the full file
		var blocks []carBlock
		for block, err := range blocksIter {
			if err != nil {
				// error reading, invalid data
				t.Skip()
			}
			blocks = append(blocks, block)
		}

		var buf bytes.Buffer
		err = writeCar(&buf, roots, func(yield func(carBlock, error) bool) {
			for _, blk := range blocks {
				if !yield(blk, nil) {
					return
				}
			}
		})
		require.NoError(t, err)

		// test if the round-trip produce a byte-equal CAR
		require.Equal(t, data, buf.Bytes())
	})
}
