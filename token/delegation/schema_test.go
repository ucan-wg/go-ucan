package delegation_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

//go:embed delegation.ipldsch
var schemaBytes []byte

func TestSchemaRoundTrip(t *testing.T) {
	t.Parallel()

	delegationJson := golden.Get(t, "new.dagjson")
	privKey := didtest.PersonaAlice.PrivKey()

	t.Run("via buffers", func(t *testing.T) {
		t.Parallel()

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()        Unseal()          EncodeDagJson()

		p1, err := delegation.FromDagJson(delegationJson)
		require.NoError(t, err)

		_, newCID, err := p1.ToSealed(privKey)
		require.NoError(t, err)

		cborBytes, id, err := p1.ToSealed(privKey)
		require.NoError(t, err)
		assert.Equal(t, envelope.CIDToBase58BTC(newCID), envelope.CIDToBase58BTC(id))

		p2, c2, err := delegation.FromSealed(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, id, c2)

		readJson, err := p2.ToDagJson(privKey)
		require.NoError(t, err)

		assert.JSONEq(t, string(delegationJson), string(readJson))
	})

	t.Run("via streaming", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(delegationJson)

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()         Unseal()         EncodeDagJson()

		p1, err := delegation.FromDagJsonReader(buf)
		require.NoError(t, err)

		_, newCID, err := p1.ToSealed(privKey)
		require.NoError(t, err)

		cborBytes := &bytes.Buffer{}
		id, err := p1.ToSealedWriter(cborBytes, privKey)
		require.NoError(t, err)
		assert.Equal(t, envelope.CIDToBase58BTC(newCID), envelope.CIDToBase58BTC(id))

		// buf = bytes.NewBuffer(cborBytes.Bytes())
		p2, c2, err := delegation.FromSealedReader(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, envelope.CIDToBase58BTC(id), envelope.CIDToBase58BTC(c2))

		readJson := &bytes.Buffer{}
		require.NoError(t, p2.ToDagJsonWriter(readJson, privKey))

		assert.JSONEq(t, string(delegationJson), readJson.String())
	})

	t.Run("fails with wrong PrivKey", func(t *testing.T) {
		t.Parallel()

		p1, err := delegation.FromDagJson(delegationJson)
		require.NoError(t, err)

		_, _, err = p1.ToSealed(didtest.PersonaBob.PrivKey())
		require.EqualError(t, err, "private key doesn't match the issuer")
	})
}

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}

func BenchmarkRoundTrip(b *testing.B) {
	delegationJson := golden.Get(b, "new.dagjson")
	privKey := didtest.PersonaAlice.PrivKey()

	b.Run("via buffers", func(b *testing.B) {
		p1, _ := delegation.FromDagJson(delegationJson)
		cborBytes, _, _ := p1.ToSealed(privKey)
		p2, _, _ := delegation.FromSealed(cborBytes)

		b.ResetTimer()

		b.Run("FromDagJson", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = delegation.FromDagJson(delegationJson)
			}
		})

		b.Run("Seal", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _, _ = p1.ToSealed(privKey)
			}
		})

		b.Run("Unseal", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _, _ = delegation.FromSealed(cborBytes)
			}
		})

		b.Run("ToDagJson", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = p2.ToDagJson(privKey)
			}
		})
	})

	b.Run("via streaming", func(b *testing.B) {
		p1, _ := delegation.FromDagJsonReader(bytes.NewReader(delegationJson))
		cborBuf := &bytes.Buffer{}
		_, _ = p1.ToSealedWriter(cborBuf, privKey)
		cborBytes := cborBuf.Bytes()
		p2, _, _ := delegation.FromSealedReader(bytes.NewReader(cborBytes))

		b.ResetTimer()

		b.Run("FromDagJsonReader", func(b *testing.B) {
			b.ReportAllocs()
			reader := bytes.NewReader(delegationJson)
			for i := 0; i < b.N; i++ {
				_, _ = reader.Seek(0, 0)
				_, _ = delegation.FromDagJsonReader(reader)
			}
		})

		b.Run("SealWriter", func(b *testing.B) {
			b.ReportAllocs()
			buf := &bytes.Buffer{}
			for i := 0; i < b.N; i++ {
				buf.Reset()
				_, _ = p1.ToSealedWriter(buf, privKey)
			}
		})

		b.Run("UnsealReader", func(b *testing.B) {
			b.ReportAllocs()
			reader := bytes.NewReader(cborBytes)
			for i := 0; i < b.N; i++ {
				_, _ = reader.Seek(0, 0)
				_, _, _ = delegation.FromSealedReader(reader)
			}
		})

		b.Run("ToDagJsonReader", func(b *testing.B) {
			b.ReportAllocs()
			buf := &bytes.Buffer{}
			for i := 0; i < b.N; i++ {
				buf.Reset()
				_ = p2.ToDagJsonWriter(buf, privKey)
			}
		})
	})
}
