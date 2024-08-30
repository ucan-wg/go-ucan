package command

import (
	"errors"
	"fmt"
	"strings"
)

const (
	separator = "/"
)

// ErrNew indicates that the wrapped error was encountered while creating
// a new Command.
var ErrNew = errors.New("failed to create Command from elems")

// ErrParse indicates that the wrapped error was encountered while
// attempting to parse a string as a Command.
var ErrParse = errors.New("failed to parse Command")

// ErrorJoin indicates that the wrapped error was encountered while
// attempting to join a new segment to a Command.
var ErrJoin = errors.New("failed to join segments to Command")

// ErrRequiresLeadingSlash is returned when a parsing a string that
// doesn't start with a [leading slash character].
//
// [leading slash character]: https://github.com/ucan-wg/spec#segment-structure
var ErrRequiresLeadingSlash = parseError("a command requires a leading slash character")

// ErrDisallowsTrailingSlash is returned when parsing a string that [ends
// with a trailing slash character].
//
// [ends with a trailing slash character]: https://github.com/ucan-wg/spec#segment-structure
var ErrDisallowsTrailingSlash = parseError("a command must not include a trailing slash")

// ErrUCANNamespaceReserved is returned to indicate that a Command's
// first segment would contain the [reserved "ucan" namespace].
//
// [reserved "ucan" namespace]: https://github.com/ucan-wg/spec#ucan-namespace
var ErrUCANNamespaceReserved = errors.New("the UCAN namespace is reserved")

// ErrRequiresLowercase is returned if a Command contains, or would contain,
// [uppercase unicode characters].
//
// [uppercase unicode characters]: https://github.com/ucan-wg/spec#segment-structure
var ErrRequiresLowercase = parseError("UCAN path segments must must not contain upper-case characters")

func parseError(msg string) error {
	return fmt.Errorf("%w: %s", ErrParse, msg)
}

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
func New(segments ...string) (*Command, error) {
	return newCommand(ErrNew, segments...)
}

func newCommand(err error, segments ...string) (*Command, error) {
	if len(segments) > 0 && segments[0] == "ucan" {
		return nil, fmt.Errorf("%w: %w", err, ErrUCANNamespaceReserved)
	}

	cmd := Command{segments}

	return &cmd, nil
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
	return newCommand(ErrParse, strings.Split(s, "/")[1:]...)
}

// [Top] is the most powerful capability.
//
// This function returns a Command that is a wildcard and therefore represents the
// most powerful abilily.  As such it should be handle with care and used
// sparingly.
//
// [Top]: https://github.com/ucan-wg/spec#-aka-top
func Top() *Command {
	cmd, _ := New()

	return cmd
}

// IsValid returns true if the provided string is a valid UCAN command.
func IsValid(s string) bool {
	_, err := Parse(s)

	return err == nil
}

// Join appends segments to the end of this command using the required
// segment separator.
func (c *Command) Join(segments ...string) (*Command, error) {
	return newCommand(ErrJoin, append(c.segments, segments...)...)
}

// Segments returns the ordered segments that comprise the Command as a
// slice of strings.
func (c *Command) Segments() []string {
	return c.segments
}

// String returns the composed representation the command.  This is also
// the required wire representation (before IPLD encoding occurs.)
func (c *Command) String() string {
	return "/" + strings.Join([]string(c.segments), "/")
}
