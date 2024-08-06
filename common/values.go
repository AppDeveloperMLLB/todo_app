package common

import (
	"context"
	"net/http"
)

type traceIdKey struct{}

func SetTraceId(ctx context.Context, traceId int) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

func GetTraceId(ctx context.Context) int {
	id := ctx.Value(traceIdKey{})
	if idInt, ok := id.(int); ok {
		return idInt
	}

	return 0
}

type userNameKey struct{}

// GetUserName gets the user name from the request context
func GetUserName(ctx context.Context) string {
	name := ctx.Value(userNameKey{})
	if nameStr, ok := name.(string); ok {
		return nameStr
	}

	return ""
}

// SetUserName sets the user name in the request context
func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)
	return req
}

type userIDKey struct{}

// GetUserID gets the user id from the request context
func GetUserID(ctx context.Context) string {
	id := ctx.Value(userIDKey{})
	if idStr, ok := id.(string); ok {
		return idStr
	}

	return ""
}

// SetUserID sets the user id in the request context
func SetUserID(req *http.Request, id string) *http.Request {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userIDKey{}, id)
	req = req.WithContext(ctx)
	return req
}
