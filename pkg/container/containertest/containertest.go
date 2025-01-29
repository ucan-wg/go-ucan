package containertest

import _ "embed"

//go:embed Base64StdPadding
var Base64StdPadding string

//go:embed Base64StdPaddingGzipped
var Base64StdPaddingGzipped string

//go:embed Base64URL
var Base64URL string

//go:embed Base64URLGzipped
var Base64URLGzipped string

//go:embed Bytes
var Bytes []byte

//go:embed BytesGzipped
var BytesGzipped []byte
