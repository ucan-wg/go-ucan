package envelope_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	_ "embed"
	"encoding/base64"
	"testing"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"github.com/ucan-wg/go-ucan/internal/token"
	"gotest.tools/v3/golden"
)

const (
	exampleSignature    = "G9EFlDm5csIZR+byd5qMFxuaN/gsZmPSeoecW2PqWW8+wYWna9zx0peX1g7mUdo4ZTLTTr8LJSxuF1JFOJR0EsjgM0c8OHuX0WpSv8U+KSNxonbZpZqO8lyI/kW4crl/k9QrWMXtyHLEOS1OD3q9SsNGsf62fk1AMH9W+D2JVBVWdWAYFVXVkXQ+RbJi21lWYc9v/JtHSJbbuCbwhRqEsTBdhcYnyfFLcgLZvR9vqM636gA3ebRjZGZJOiAvxwdTOzlVxtw/552pAx8Od3hRGc5xdG5jGu2/OwIn9UMoXPQl7pMUYqk1nfqN3C7kDelIaQlgoAGyfssepB1tMRH/KA"
	exampleTag          = "ucan/example@v1.0.0-rc.1"
	exampleVarsigHeader = "NIUkEoACcQ"
	invalidSignature    = "a5BocvMSlifrDzWN7MQpDZ4cEciwe+b9twdQ7d5EZ/LlW3w1VIjk34ci8LqmzMCMwqJsoBqevArUMNS86RrDOLZEl+71+nSf1GJ9fK/E2o7ONSPTQt1wALH1xhJ4S/h5o8v0sWP/PWBvolSfMpro9lN1xCi9zC4iuFmizqdjOd3Ba3txHD5DGAculWBiob3N1mjkXZPbQYEQteCoLwSNDCmmHCE7VpRUkoi832N7UVHlu1FFucENB31qBWZQ+JTj8/oV56Do+LbhrDDiabNkTxulwQ7u+hdKA30vA6FWaA6QW+UE2/mCEKM5wvVAohLPZsapGXP6LoEcbBM3O758dx"
	priCfg              = "CAASqAkwggSkAgEAAoIBAQDq39Aou82MEteoEz+iKpu7zwJc0dZfomAfB4Zpnl8+WhUOhZyveHDD9lr/UCc/fcN5ufeyZxutDRvIXcmUGG5DNTVRZ10ywT/wN8KO+x/hZ5QIxBAsCFukcyHbAPseLYpAK0J0HNnQhtF6cQcQkuThCnZH/Ofj42d7snuztbBUxwjButvHYHiWwolcJUeb99HCpGwtJYjkp004roFBqjkLayP4AWHrnW4mtCY0rw86gRCT60N1XBZ9zXKw+LJeuQg3RUgZqBL6hvVIs1LAY5ie0LSXVkdjg9bmV7j5SKJBgk9ABoKvt35/KSWA5HW/6g3y/UITCD2DQDrTFv8xzDIvAgMBAAECggEAFoGr6LtWTv3fPHPbvSZoFe8YQty4tiFRJKgL8UMDzW3EZsfW49metKh+v8hmemcKvDddzPKkbEi9SM3z6wUMS9Rlb4+AFsT94370Xc8ilu7d+JkRE6cZYQDHVb0aUyH6BXwfuhCprpm8qQb7rlLlK8tc2jkZ33SDDg9kWywl4XmiDabKm0fOJd68KGuO5FNpCfpipqG3ok/FYuKlSqpCz+7QH2p3z2eVGTa/uIbDMNxkFBoIuhEJT43eDR2elPOrSL9+AYgBPVrzcJoRRxZFVvDDQ+RIwI/A62DvZ3IpFEyzEk2VZwWpLWYKnAUElcSegxx22K7S4BAaWjL86I6cAQKBgQDr8Y1NLYKYffHJAkmWE/ssGcih5C06gOo9WInBU8ZroY4LAhzKLstS0yKsM0OtRSFhZsmCdR3M/IHO94c3KsXu+KtA7r9AG+58LMpmob1mvyGsXIowHFMAd95s3FDd/HOE10tMyrIE1c7eWLfII+s6yGo8MS0WHXOSFBlioukGbwKBgQD+1v3vC+iUNP0FMhVxnhGgONyUb0X+AEw9GOLeCpsugZqRXnharYSiTjGPjwPaT2YBVkyaraoX6VwK+ys3RCngUg4s9IeHUYsR6Xa0oheAlME0ZDtuzi5+lDo+Zpsg8vepgx6v//bKvUcJb+9YDcKlMfQgFnSb3bAwUN9Ru79wQQKBgD7Z/9wZTXq1whzbwSJ7fCNJUwrdL7cv9DYXScr4OBkf1ijUjTrGsF8F42yf011q1vONYAyiiie69BFgGuL1P/jiwSvw7X10c1kczWX9m+is7ZlupVkfknTDebriDaC0yUkP2P1B2Z40HoFYfMyR1O25yaLzLqF/gvPc6s49u3l9AoGBAOZ9d2EdKTfbEToAyYpgyFpc84zBc9G/XTUpbBAeEasnh7CRfFOvezX9eS/5zydGBuGQt2pzRlOoOhqof7bVzPZZ4P5iEK6gXyNNQJMxxAYFBRYozeRzUXQlBuTnksljWAMWV8whu4o1VanAdv7yOymEm+Ply4QqJzAcBU/8erLBAoGBAJ6120n/Q8B7kNW/tZvJ4Xi0u/kSEodNQ9TpF44SB32bn/aWwfe7qFS1x+3omO7XWcx3FLUUPhITQjmQcbNa2yWY0UZoqnzkHhDmJeG2PUILEMCHSLCKQHS+PNJxqEWvwQe2mX/gJIjf14U9983hgLnOL7gH9sVYf9M9yA8NVlem"
)

func TestNew(t *testing.T) {
	t.Parallel()

	exampleSignature, err := base64.RawStdEncoding.DecodeString(exampleSignature)
	require.NoError(t, err)
	varsigHeader, err := base64.RawStdEncoding.DecodeString(exampleVarsigHeader)
	require.NoError(t, err)

	env := exampleEnvelope(t)
	assert.Equal(t, exampleSignature, env.Signature())
	assert.Equal(t, varsigHeader, env.VarsigHeader())

	tkn := env.Token()
	assert.IsType(t, (*token.Token)(nil), tkn)
	// TODO
}

func TestWrap(t *testing.T) {
	t.Parallel()

	envNode, err := envelope.Wrap(rsaPrivateKey(t), exampleToken(t), exampleTag)
	assert.NoError(t, err)
	assert.NotNil(t, envNode)

	buf := &bytes.Buffer{}
	require.NoError(t, dagcbor.Encode(envNode, buf))

	golden.AssertBytes(t, buf.Bytes(), "example.cbor")

	// TODO: use golden file
}

func TestEnvelope_Wrap(t *testing.T) {
	t.Parallel()

	env := exampleEnvelope(t)

	envNode, err := env.Wrap()
	require.NoError(t, err)

	buf := &bytes.Buffer{}
	require.NoError(t, dagjson.Encode(envNode, buf))

	golden.AssertBytes(t, buf.Bytes(), "example.json")

	t.Log(buf.String())

	env1, err := envelope.Unwrap(envNode, exampleTag)
	require.NoError(t, err)
	assert.NotNil(t, env1)

	t.Log(string(env1.Signature()))
	assert.Equal(t, env.Signature(), env1.Signature())
	assert.Equal(t, env.VarsigHeader(), env1.VarsigHeader())
	assert.Equal(t, env.Token(), env1.Token())

	// t.Fail()
}

func TestEnvelope_Verify(t *testing.T) {
	t.Parallel()

	t.Run("true with correct public key", func(t *testing.T) {
		t.Parallel()

		env := exampleEnvelope(t)
		ok, err := env.Verify(rsaPublicKey(t))
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("false with wrong public key", func(t *testing.T) {
		t.Parallel()

		_, pub, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
		require.NoError(t, err)

		env := exampleEnvelope(t)
		ok, err := env.Verify(pub)
		assert.ErrorIs(t, err, rsa.ErrVerification)
		assert.False(t, ok)
	})
}

func exampleEnvelope(t *testing.T) *envelope.Envelope {
	t.Helper()

	env, err := envelope.New(rsaPrivateKey(t), exampleToken(t), exampleTag)
	require.NoError(t, err)

	t.Log("exampleEnvelope.Signature", base64.RawStdEncoding.EncodeToString(env.Signature()))

	return env
}

func exampleToken(t *testing.T) *token.Token {
	t.Helper()

	id, err := did.Parse("did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK")
	require.NoError(t, err)
	_ = id // TODO:

	tkn, err := token.New() // TODO: fields
	require.NoError(t, err)

	return tkn
}

func rsaPrivateKey(t *testing.T) crypto.PrivKey {
	t.Helper()

	priEnc, err := crypto.ConfigDecodeKey(priCfg)
	require.NoError(t, err)
	pri, err := crypto.UnmarshalPrivateKey(priEnc)
	require.NoError(t, err)

	return pri
}

func rsaPublicKey(t *testing.T) crypto.PubKey {
	t.Helper()

	return rsaPrivateKey(t).GetPublic()
}
