package bearer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ucan-wg/go-ucan/pkg/container"
)

var ErrNoUcan = fmt.Errorf("no ucan")
var ErrContainerMalformed = fmt.Errorf("malformed container")

// ExtractBearerContainer extract a full UCAN container from an HTTP "Authorization: Bearer" header.
// It can return:
// - ErrNoUcan if no such HTTP header is found
// - ErrContainerMalformed if the container or token can't be decoded or if a token is invalid (bad signature ...)
func ExtractBearerContainer(h http.Header) (container.Reader, error) {
	header := h.Get("Authorization")
	if header == "" {
		return nil, ErrNoUcan
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return nil, ErrNoUcan
	}

	// skip prefix
	reader := strings.NewReader(header[len("Bearer "):])

	// read container from any supported format
	ctn, err := container.FromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrContainerMalformed, err)
	}

	return ctn, nil
}

// AddBearerContainer adds the given UCAN container into an `Authorization: Bearer` HTTP header.
//
// You should use this with HTTP/2 or HTTP/3, as both of those already compress their header.
func AddBearerContainer(h http.Header, container container.Writer) error {
	str, err := container.ToBase64StdPadding()
	if err != nil {
		return err
	}

	h.Set("Authorization", fmt.Sprintf("Bearer %s", str))
	return nil
}

// AddBearerContainerCompressed adds the given UCAN container into an `Authorization: Bearer` HTTP header.
//
// You should use this with HTTP/1, as it doesn't compress its headers.
func AddBearerContainerCompressed(h http.Header, container container.Writer) error {
	str, err := container.ToBase64StdPaddingGzipped()
	if err != nil {
		return err
	}

	h.Set("Authorization", fmt.Sprintf("Bearer %s", str))
	return nil
}
