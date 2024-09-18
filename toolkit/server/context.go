package server

import "context"

type contextKey string

var ctxUserId = contextKey("userId")
var ctxProjectId = contextKey("projectId")

// ContextGetUserId return the UserId stored in the context, if it exists.
func ContextGetUserId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(ctxUserId).(string)
	return val, ok
}

func addUserIdToContext(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, ctxUserId, userId)
}

// ContextGetProjectId return the ProjectID stored in the context, if it exists.
func ContextGetProjectId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(ctxProjectId).(string)
	return val, ok
}

func addProjectIdToContext(ctx context.Context, projectId string) context.Context {
	return context.WithValue(ctx, ctxProjectId, projectId)
}
