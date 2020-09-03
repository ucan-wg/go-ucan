package ucan

import (
	"context"
)

// CtxKey defines a distinct type for context keys used by the access
// package
type CtxKey string

// UCANCtxKey is the key for adding an access UCAN to a context.Context
const UCANCtxKey CtxKey = "UCAN"

// CtxWithUCAN adds a UCAN value to a context
func CtxWithUCAN(ctx context.Context, t UCAN) context.Context {
	return context.WithValue(ctx, UCANCtxKey, t)
}

// UCANFromCtx extracts a Dataset reference from a given
// context if one is set, returning nil otherwise
func UCANFromCtx(ctx context.Context) *UCAN {
	iface := ctx.Value(UCANCtxKey)
	if ref, ok := iface.(*UCAN); ok {
		return ref
	}
	return nil
}
