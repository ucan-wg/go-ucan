package delegation

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
)

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}
