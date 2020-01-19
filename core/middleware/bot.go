package middleware

import "context"

const BotContextKey = MiddlewareKey("bot")

type BotData struct {
	Name string
}

func WithBot(ctx context.Context) context.Context {
	return context.WithValue(ctx, BotContextKey, BotData{
		Name: "Geralt",
	})
}
