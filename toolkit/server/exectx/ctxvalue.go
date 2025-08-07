package exectx

import "context"

type ctxKey struct{}

// AddUcanCtxToContext insert a UcanCtx into a go context.
func AddUcanCtxToContext(ctx context.Context, ucanCtx *UcanCtx) context.Context {
	return context.WithValue(ctx, ctxKey{}, ucanCtx)
}

// FromContext retrieve a UcanCtx from a go context.
func FromContext(ctx context.Context) (*UcanCtx, bool) {
	ucanCtx, ok := ctx.Value(ctxKey{}).(*UcanCtx)
	return ucanCtx, ok
}
