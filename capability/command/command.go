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
type Command struct {
	segments []string
}

// New creates a validated command from the provided list of segment
// strings.  An error is returned if an invalid Command would be
// formed
func New(segments ...string) *Command {
	return &Command{segments: segments}
}

// Parse verifies that the provided string contains the required
// [segment structure] and, if valid, returns the resulting
// Command.
//
// [segment structure]: https://github.com/ucan-wg/spec#segment-structure
func Parse(s string) (*Command, error) {
	if !strings.HasPrefix(s, "/") {
		return nil, ErrRequiresLeadingSlash
	}

	if len(s) > 1 && strings.HasSuffix(s, "/") {
		return nil, ErrDisallowsTrailingSlash
	}

	if s != strings.ToLower(s) {
		return nil, ErrRequiresLowercase
	}

	// The leading slash will result in the first element from strings.Split
	// being an empty string which is removed as strings.Join will ignore it.
	return &Command{strings.Split(s, "/")[1:]}, nil
}

// MustParse is the same as Parse, but panic() if the parsing fail.
func MustParse(s string) *Command {
	c, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

// [Top] is the most powerful capability.
//
// This function returns a Command that is a wildcard and therefore represents the
// most powerful abilily.  As such it should be handle with care and used
// sparingly.
//
// [Top]: https://github.com/ucan-wg/spec#-aka-top
func Top() *Command {
	return New()
}

// IsValid returns true if the provided string is a valid UCAN command.
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}

// Join appends segments to the end of this command using the required
// segment separator.
func (c *Command) Join(segments ...string) *Command {
	return &Command{append(c.segments, segments...)}
}

// Segments returns the ordered segments that comprise the Command as a
// slice of strings.
func (c *Command) Segments() []string {
	return c.segments
}

// String returns the composed representation the command.  This is also
// the required wire representation (before IPLD encoding occurs.)
func (c *Command) String() string {
	return "/" + strings.Join(c.segments, "/")
}
