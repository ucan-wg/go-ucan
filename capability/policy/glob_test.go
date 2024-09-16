package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleGlobMatch(t *testing.T) {
	tests := []struct {
		pattern string
		str     string
		matches bool
	}{
		// Basic matching
		{"*", "anything", true},
		{"a*", "abc", true},
		{"*c", "abc", true},
		{"a*c", "abc", true},
		{"a*c", "abxc", true},
		{"a*c", "ac", true},
		{"a*c", "a", false},
		{"a*c", "ab", false},

		// Escaped characters
		{"a\\*c", "a*c", true},
		{"a\\*c", "abc", false},

		// Mixed wildcards and literals
		{"a*b*c", "abc", true},
		{"a*b*c", "aXbYc", true},
		{"a*b*c", "aXbY", false},
		{"a*b*c", "abYc", true},
		{"a*b*c", "aXbc", true},
		{"a*b*c", "aXbYcZ", false},

		// Edge cases
		{"", "", true},
		{"", "a", false},
		{"*", "", true},
		{"*", "a", true},
		{"\\*", "*", true},
		{"\\*", "a", false},

		// Specified test cases
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob, Carol.", true},
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob, Dan, Erin, Carol.", true},
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob  , Carol.", true},
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob*, Carol.", true},
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob, Carol", false},
		{"Alice\\*, Bob*, Carol.", "Alice*, Bob*, Carol!", false},
		{"Alice\\*, Bob*, Carol.", "Alice, Bob, Carol.", false},
		{"Alice\\*, Bob*, Carol.", "Alice Cooper, Bob, Carol.", false},
		{"Alice\\*, Bob*, Carol.", " Alice*, Bob, Carol. ", false},
	}

	for _, tt := range tests {
		t.Run(tt.pattern+"_"+tt.str, func(t *testing.T) {
			assert.Equal(t, tt.matches, globMatch(tt.pattern, tt.str))
		})
	}
}
