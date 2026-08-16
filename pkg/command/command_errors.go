package command

import "fmt"

// ErrRequiresLeadingSlash is returned when a parsing a string that
// doesn't start with a [leading slash character].
//
// [leading slash character]: https://github.com/ucan-wg/spec#segment-structure
var ErrRequiresLeadingSlash = fmt.Errorf("a command requires a leading slash character")

// ErrDisallowsTrailingSlash is returned when parsing a string that [ends
// with a trailing slash character].
//
// [ends with a trailing slash character]: https://github.com/ucan-wg/spec#segment-structure
var ErrDisallowsTrailingSlash = fmt.Errorf("a command must not include a trailing slash")

// ErrRequiresLowercase is returned if a Command contains, or would contain,
// [uppercase unicode characters].
//
// [uppercase unicode characters]: https://github.com/ucan-wg/spec#segment-structure
var ErrRequiresLowercase = fmt.Errorf("UCAN path segments must must not contain upper-case characters")

// ErrEmptySegment is returned when parsing a string in which two [segment
// separators] are adjacent, as in "/crud//create".
//
// [segment separators]: https://github.com/ucan-wg/spec#segment-structure
var ErrEmptySegment = fmt.Errorf("a command must not contain empty segments")
