package middleware

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const BotContextKey = MiddlewareKey("bot")

type BotData struct {
	Name string
}

func WithBot(ctx context.Context, db *mongo.Database) context.Context {
	return context.WithValue(ctx, BotContextKey, BotData{
		Name: "Geralt",
	})
}
