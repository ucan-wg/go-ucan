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

	err = writeCar(buf, roots, func(yield func(carBlock) bool) {
		for _, blk := range blks {
			if !yield(blk) {
				return
			}
		}
	})
	require.NoError(t, err)

	// Bytes equal after the round-trip
	require.Equal(t, original, buf.Bytes())
}

func FuzzCarRead(f *testing.F) {
	example, err := os.ReadFile("testdata/sample-v1.car")
	require.NoError(f, err)

	f.Add(example)

	f.Fuzz(func(t *testing.T, data []byte) {
		_, _, _ = readCar(bytes.NewReader(data))
		// only looking for panics
	})
}
