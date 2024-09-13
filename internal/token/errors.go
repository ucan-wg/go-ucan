package token

import "errors"

var ErrFailedSchemaLoad = errors.New("failed to load IPLD Schema")

var ErrNoSchemaType = errors.New("schema does not contain type")

var ErrNodeNotToken = errors.New("IPLD node is not a Token")

var ErrMissingRequiredDID = errors.New("a required DID is missing")
