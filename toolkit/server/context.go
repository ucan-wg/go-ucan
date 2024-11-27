package server

import "context"

type ctxUserIdKey struct{}

// ContextGetUserId return the UserId stored in the context, if it exists.
func ContextGetUserId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(ctxUserIdKey{}).(string)
	return val, ok
}

func addUserIdToContext(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, ctxUserIdKey{}, userId)
}

type ctxProjectIdKey struct{}

// ContextGetProjectId return the ProjectID stored in the context, if it exists.
func ContextGetProjectId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(ctxProjectIdKey{}).(string)
	return val, ok
}

func addProjectIdToContext(ctx context.Context, projectId string) context.Context {
	return context.WithValue(ctx, ctxProjectIdKey{}, projectId)
}
