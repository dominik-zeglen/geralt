package middleware

import "context"

const UserContextKey = MiddlewareKey("user")

type UserData struct {
	Name string
}

func WithUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, UserContextKey, UserData{
		Name: "Dominik",
	})
}
