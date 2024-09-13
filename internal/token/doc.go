// Package token provides a generic model of the [TokenPayload] required
// within an Envelope.
//
// # Field requirements
//
// While the Token object represents the wire format of both a UCAN
// Delegation token and a UCAN Invocation token, the delegation and
// invocation packages are, respectively, responsible for making sure
// required fields are included when creating a new Token or when
// validating the contents of an Envelope as it's received from
// another party.  The following table shows the current (as of
// 2024-09-11) relationship between optional and nullable fields in
// the delegation and invocation views and the payload model:
//
//	| Name  |     Delegation      |     Invocation      |  Token   |
//	|       | Required | Nullable | Required | Nullable |          |
//	| ----- | -------- | -------- | -------- | -------- | -------- |
//	| iss   | Yes      | No       | Yes      | No       |          |
//	| aud   | Yes      | No       | No       | N/A      | Optional |
//	| sub   | Yes      | Yes      | Yes      | No       | Nullable |
//	| cmd   | Yes      | No       | Yes      | No       |          |
//	| pol   | Yes      | No       | X        | X        | Optional |
//	| nonce | Yes      | No       | No       | N/A      | Optional |
//	| meta  | No       | N/A      | No       | N/A      | Optional |
//	| nbf   | No       | N/A      | X        | X        | Optional |
//	| exp   | Yes      | Yes      | Yes      | Yes      |          |
//	| args  | X        | X        | Yes      | No       | Optional |
//	| prf   | X        | X        | Yes      | No       | Optional |
//	| iat   | X        | X        | No       | N/A      | Optional |
//	| cause | X        | X        | No       | N/A      | Optional |
//
// [TokenPayload]: https://github.com/ucan-wg/spec?tab=readme-ov-file#envelope
package token
