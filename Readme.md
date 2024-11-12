<div align="center">
  <a href="https://github.com/ucan-wg/go-ucan" target="_blank">
    <img src="https://raw.githubusercontent.com/ucan-wg/go-ucan/v1/assets/logo.png" alt="go-ucan Logo" height="250"></img>
  </a>

  <h1 align="center">go-ucan</h1>

  <p>
    <img src="https://img.shields.io/badge/UCAN-v1.0.0--rc.1-blue" alt="UCAN v1.0.0-rc.1">
    <a href="https://github.com/ucan-wg/go-ucan/tags">
        <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/ucan-wg/go-ucan">
    </a>
    <a href="https://github.com/ucan-wg/go-ucan/actions?query=">
      <img src="https://github.com/ucan-wg/go-ucan/actions/workflows/gotest.yml/badge.svg" alt="Build Status">
    </a>
    <a href="https://ucan-wg.github.io/go-ucan/dev/bench/">
        <img alt="Go benchmarks" src="https://img.shields.io/badge/Benchmarks-go-blue">
    </a>
    <a href="https://github.com/ucan-wg/go-ucan/blob/v1/LICENSE.md">
        <img alt="Apache 2.0 + MIT License" src="https://img.shields.io/badge/License-Apache--2.0+MIT-green">
    </a>
    <a href="https://pkg.go.dev/github.com/ucan-wg/go-ucan">
      <img src="https://img.shields.io/badge/Docs-godoc-blue" alt="Docs">
    </a>
    <a href="https://discord.gg/JSyFG6XgVM">
      <img src="https://img.shields.io/static/v1?label=Discord&message=join%20us!&color=mediumslateblue" alt="Discord">
    </a>
  </p>
</div>

This is a go library to help the next generation of web and decentralized applications make use
of UCANs in their authorization flows.

User Controlled Authorization Networks (UCANs) are a way of doing authorization where users are fully in control. OAuth is designed for a centralized world, UCAN is the distributed user controlled version.

## Resources

### Specifications

The UCAN specification is separated in multiple sub-spec:
- [Main specification](https://github.com/ucan-wg/spec)
- [Delegation](https://github.com/ucan-wg/delegation/tree/v1_ipld)
- [Invocation](https://github.com/ucan-wg/invocation)

Not implemented yet:
- [Revocation](https://github.com/ucan-wg/revocation/tree/first-draft)
- [Promise](https://github.com/ucan-wg/promise/tree/v1-rc1)

### Talks

- [Decentralizing Auth, and UCAN Too - Brooklyn Zelenka (2023)](https://www.youtube.com/watch?v=MuHfrqw9gQA)
- [What's New in UCAN 1.0 - Brooklyn Zelenka (2024)](https://www.youtube.com/watch?v=-uohQzZcwF4)

## Status

`go-ucan` currently support the required parts of the UCAN specification: the main specification, delegation and invocation.

Besides that, `go-ucan` also includes:
- a simplified [DID](https://www.w3.org/TR/did-core/) and [did-key](https://w3c-ccg.github.io/did-method-key/) implementation
- a [token container](https://github.com/ucan-wg/go-ucan/tree/v1/pkg/container) with CBOR and CAR format, to package and carry tokens together
- support for encrypted values in token's metadata

## Getting Help

For usage questions, usecases, or issues reach out to us in our `go-ucan`
[Discord channel](https://discord.gg/3EHEQ6M8BC).

We would be happy to try to answer your question or try opening a new issue on
Github.

## UCAN Gopher

Artwork by [Bruno Monts](https://www.instagram.com/bruno_monts). Thank you [Renee French](http://reneefrench.blogspot.com/) for creating the [Go Gopher](https://blog.golang.org/gopher)

## License

This project is licensed under the double license [Apache 2.0 + MIT](https://github.com/ucan-wg/go-ucan/blob/v1/LICENSE.md).
