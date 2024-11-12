package command

import (
	"fmt"
	"strings"
)

const separator = "/"

var _ fmt.Stringer = (*Command)(nil)

// Command is a concrete messages (a "verb") that MUST be unambiguously
// interpretable by the Subject of a UCAN.
//
// A [Command] is composed of a leading slash which is optionally followed
// by one or more slash-separated Segments of lowercase characters.
//
// [Command]: https://github.com/ucan-wg/spec#command
type Command string

// New creates a validated command from the provided list of segment strings.
// An error is returned if an invalid Command would be formed
func New(segments ...string) Command {
	return Top().Join(segments...)
}

// Parse verifies that the provided string contains the required
// [segment structure] and, if valid, returns the resulting
// Command.
//
// [segment structure]: https://github.com/ucan-wg/spec#segment-structure
func Parse(s string) (Command, error) {
	if !strings.HasPrefix(s, "/") {
		return "", ErrRequiresLeadingSlash
	}

	if len(s) > 1 && strings.HasSuffix(s, "/") {
		return "", ErrDisallowsTrailingSlash
	}

	if s != strings.ToLower(s) {
		return "", ErrRequiresLowercase
	}

	// The leading slash will result in the first element from strings.Split
	// being an empty string which is removed as strings.Join will ignore it.
	return Command(s), nil
}

// MustParse is the same as Parse, but panic() if the parsing fail.
func MustParse(s string) Command {
	c, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

// Top is the most powerful capability.
//
// This function returns a Command that is a wildcard and therefore represents the
// most powerful ability. As such, it should be handled with care and used sparingly.
//
// [Top]: https://github.com/ucan-wg/spec#-aka-top
func Top() Command {
	return Command(separator)
}

// IsValid returns true if the provided string is a valid UCAN command.
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}

// Join appends segments to the end of this command using the required
// segment separator.
func (c Command) Join(segments ...string) Command {
	size := 0
	for _, s := range segments {
		size += len(s)
	}
	if size == 0 {
		return c
	}
	buf := make([]byte, 0, len(c)+size+len(segments))
	buf = append(buf, []byte(c)...)
	for _, s := range segments {
		if s != "" {
			if len(buf) > 1 {
				buf = append(buf, separator...)
			}
			buf = append(buf, []byte(s)...)
		}
	}
	return Command(buf)
}

// Segments returns the ordered segments that comprise the Command as a
// slice of strings.
func (c Command) Segments() []string {
	if c == separator {
		return nil
	}
	return strings.Split(string(c), separator)[1:]
}

// Covers returns true if the command is identical or a parent of the given other command.
func (c Command) Covers(other Command) bool {
	// fast-path, equivalent to the code below (verified with fuzzing)
	if !strings.HasPrefix(string(other), string(c)) {
		return false
	}
	return c == separator || len(c) == len(other) || other[len(c)] == separator[0]

	/* -------

	otherSegments := other.Segments()
	if len(otherSegments) < len(c.Segments()) {
		return false
	}
	for i, s := range c.Segments() {
		if otherSegments[i] != s {
			return false
		}
	}
	return true

	*/
}

// String returns the composed representation the command.  This is also
// the required wire representation (before IPLD encoding occurs.)
func (c Command) String() string {
	return string(c)
}
