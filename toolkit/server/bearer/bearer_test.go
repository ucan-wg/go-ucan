package bearer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/container/containertest"
)

func TestHTTPBearer(t *testing.T) {
	for _, fn := range []func(h http.Header, container container.Writer) error{
		AddBearerContainer,
		AddBearerContainerCompressed,
	} {
		r, err := http.NewRequest(http.MethodPost, "/foo/bar", nil)
		require.NoError(t, err)

		cont, err := container.FromBytes(containertest.Bytes)
		require.NoError(t, err)

		err = fn(r.Header, cont.ToWriter())
		require.NoError(t, err)

		contRead, err := ExtractBearerContainer(r.Header)
		require.NoError(t, err)

		require.NotEmpty(t, contRead)
		require.Equal(t, cont, contRead)
	}
}
