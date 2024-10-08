# HTTP Server middleware example

This package shows an example of an HTTP middleware with the following features:
- decoding a UCAN delegation token from the `Authorization: Bearer` HTTP Header, according to the PriorityConnect Step0 way (which is not the fully correct UCAN way)
- minimal verifications, no policies or args
- retrieval of values passed as token metadata, and insertion in the go context