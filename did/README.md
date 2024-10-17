## did

### Testing

The test suite for this package includes test vectors provided by the
authors of the [`did:key` method specification](https://w3c-ccg.github.io/did-method-key/).
Some of these tests provide the public key associated with a `did:key`
as JWKs and an extra (test-only) dependency has been added to unmarshal
the JWK into a Go `struct`.  Support for the `secp256k1` encryption
algorithm is experimental (but stable in my experience) and requires the
addition of the following build tag to properly run:

```
// go:build jwx_es256k
```

WARNING: These tests will not run by default!

To include these tests from the CLI, execute the following command:

```
go test -v ./did -tags jwx_es256k
```

It should also be possible to configure your IDE to run these tests.  For
instance, in Codium, add the following JSON snippet to your local project
configuration:

```
"go.testTags": "jwx_es256k",
```