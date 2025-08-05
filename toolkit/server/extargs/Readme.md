## Motivations

UCAN is normally a pure RPC construct, when the entirety of the request's parameters is part of the invocation, in the form of `args`. Those `args` are evaluated against the delegation's [policy](https://github.com/ucan-wg/delegation/tree/v1_ipld?tab=readme-ov-file#policy) to determine if the request is allowed or not, then the request handling happens purely based on those args and the `command`. In that setup, the service would have a single entry point.

Unfortunately, we live in a world of REST APIs, or JSON-RPC. Some adaptations or concessions need to be made.

In this package, we cross the chasm of the pure UCAN world into our practical needs. This can, however, be done in a reasonable way.

## Example

Below is an example of `args` in Dag-Json format, where the values are recomposed server-side from the HTTP request:

```json
{
  "http": {
    "scheme": "https",
    "method": "POST",
    "host": "example.com",
    "path": ""
  },
  "custom": {
    "foo": "bar"
  }
}
```
Those `args` can be evaluated against a delegation's policy, for example:
```json
{
  "cmd": "/foo/bar",
  "pol": [
    ["==", ".http.host", "example.com"],
    ["like", ".custom.foo", "ba*"]
  ]
}
```

## Security implications

UCAN essentially aims for perfect security. By having external args, we break that security perimeter, and we now need to arbitrate between security and practicality.

First, what are we breaking exactly? Normally, the invocation has all the parameters and is signed by the invoker. This means that an attacker cannot intercept (MITM) a request and change the parameters when relaying it to the server. As they are signed, that would make the request invalid.

If we have external args, now an attacker can intercept the request, change it, and pretend to be that person doing other things than intended. **That may of may not be an actual problem, depending on the situation.**  

There is a way to get around that, and have the best of both worlds, but **it comes with a client side complexity**: we can hash the external values and put them into the invocation's `args`. For example:

```json
{
  "http": "zQmSnuWmxptJZdLJpKRarxBMS2Ju2oANVrgbr2xWbie9b2D",
  "custom": "zQmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"
}
```

When the server receives the request, it can now reconstruct the values from the HTTP request, verify the hash and replace it with the real values for evaluation against the policies. **We are now back to the better security model**, but at the price of client-side complexity.

## API design

Therefore, the server-side logic is made to have this hashing optional:
- if present, the server honors the hash and enforces the security
- the client can opt out of passing that hash, and won't benefit from the enforced security

The particular hash selected is SHA2-256 of the DAG-CBOR encoded argument, expressed in the form of a Multihash in raw bytes.
The arguments being hashed are the complete map of values, including the root key being replaced (for example `http` or `custom` here). 