# UCAN container Specification v0.1.0

## Editors

* [Michael Muré], [Consensys]

## Authors

* [Michael Muré], [Consensys]
* [Hugo Dias]

## Language

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [BCP 14] when, and only when, they appear in all capitals, as shown here.

# 0 Abstract

[User-Controlled Authorization Network (UCAN)][UCAN] is a trustless, secure, local-first, user-originated authorization and revocation scheme. This document describes a container format for transmitting one or more UCAN tokens as bytes, regardless of the transport.

# 1 Introduction

The UCAN spec itself is transport agnostic. This specification describes how to transfer one or more [UCAN] tokens bundled together, regardless of the transport.

# 2 Container format

## 2.1 Inner structure

UCAN tokens, regardless of their kind ([Delegation], [Invocation], [Revocation], [Promise]) MUST be first signed and serialized into bytes according to their respective specification. As the token's CID is not part of the serialized container, any CID returned by this operation is to be ignored. 

All the tokens' bytes MUST be assembled in a [CBOR] array, which is then inserted as the value under the `ctn-v1` string key, in a CBOR map. The ordering of tokens in the array MUST NOT matter. For clarity, the CBOR shape is given below:  

```json
{
  "ctn-v1": [
    <token1 bytes>,
    <token2 bytes>,
    <token3 bytes>,
  ]
}
```

## 2.2 Serialisation

To serialize the container into bytes, the inner CBOR structure MUST be serialized into bytes according to the CBOR specification. The resulting bytes MAY be compressed by a supported algorithm, then MAY be encoded with a supported base encoding.

The following compression algorithm are REQUIRED to be supported:
- [GZIP]

The following base encoding combination are REQUIRED to be supported:
- base64, standard alphabet, padding
- base64, URL alphabet, no padding

The CBOR bytes MUST be prepended by a single byte header to indicate the selection combination of base encoding and compression. This header value MUST be set according to the following table:

| Header as hex | Header as ASCII | Base encoding           | Compression    |
|---------------|-----------------|-------------------------|----------------|
| 0x40          | @               | raw bytes               | no compression | 
| 0x42          | B               | base64 std padding      | no compression | 
| 0x43          | C               | base64 url (no padding) | no compression | 
| 0x4D          | M               | raw bytes               | gzip           | 
| 0x4F          | O               | base64 std padding      | gzip           | 
| 0x50          | P               | base64 url (no padding) | gzip           | 

# 3 FAQ

## 3.1 Why not include the UCAN CIDs?

Several attacks are possible if UCAN tokens aren't validated. If CIDs aren't validated, at least two attacks are possible: [privilege escalation] and [cache poisoning], as UCAN delegation proofs depends on a correct hash-linked structure.

By not including the CID in the container, the recipient is forced to hash (and thus validate) the CIDs for each token. If presented with a claimed CID paired with the token bytes, implementers could ignore CID validation, breaking a core part of the proof chain security model. Hash functions are very fast on a couple kilobytes of data so the overhead is still very low. It also reduces significantly the size of the container.

## 3.2 Why compress? Why not always compress?

Compression is a relatively demanding operation. As such, using it is a tradeoff between size on the wire and CPU/memory usage, both when writing and reading a container. The transport itself can make compression worthwhile or note: for example, HTTP/2 and HTTP/3 headers are already compressed, but HTTP/1 headers are not. This being highly contextual, the choice is left to the final implementer.

# 4 Implementation recommendations

## 4.1 Dissociate reader and writer

While it tempting to write a single implementation to read and write a container, it is RECOMMENDED to separate the implementation into a reader and a writer. The writer can simply accept arbitrary tokens as bytes, while the reader provide a read-only view with convenient access functions.

# 5 Acknowledgments

Many thanks to all the [Fission] team and in particular to [Brooklyn Zelenka] for creating and pushing [UCAN] and other critical pieces like [WNFS], and generally being awesome and supportive people.

<!-- External Links -->

[BCP 14]: https://www.rfc-editor.org/info/bcp14
[Brooklyn Zelenka]: https://github.com/expede
[CBOR]: https://www.rfc-editor.org/rfc/rfc8949.html
[Consensys]: https://consensys.io/
[Delegation]: https://github.com/ucan-wg/delegation/tree/v1_ipld
[Fission]: https://fission.codes
[GZIP]: https://datatracker.ietf.org/doc/html/rfc1952
[Hugo Dias]: https://github.com/hugomrdias
[Invocation]: https://github.com/ucan-wg/invocation
[Michael Muré]: https://github.com/MichaelMure/
[Promise]: https://github.com/ucan-wg/promise/tree/v1-rc1
[Revocation]: https://github.com/ucan-wg/revocation/tree/first-draft
[UCAN]: https://github.com/ucan-wg/spec
[WNFS]: https://github.com/wnfs-wg
[cache poisoning]: https://en.wikipedia.org/wiki/Cache_poisoning
[privilede escalation]: https://en.wikipedia.org/wiki/Privilege_escalation
