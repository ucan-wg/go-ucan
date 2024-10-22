package selector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenizeUTF8(t *testing.T) {
	t.Run("simple UTF-8", func(t *testing.T) {
		str := ".こんにちは[0]"
		expected := []string{".", "こんにちは", "[0]"}
		actual := tokenize(str)
		require.Equal(t, expected, actual)
	})

	t.Run("UTF-8 with quotes", func(t *testing.T) {
		str := ".こんにちは[\"привет\"]"
		expected := []string{".", "こんにちは", "[\"привет\"]"}
		actual := tokenize(str)
		require.Equal(t, expected, actual)
	})

	t.Run("UTF-8 with escaped quotes", func(t *testing.T) {
		str := ".こんにちは[\"привет \\\"мир\\\"\"]"
		expected := []string{".", "こんにちは", "[\"привет \\\"мир\\\"\"]"}
		actual := tokenize(str)
		require.Equal(t, expected, actual)
	})
}
