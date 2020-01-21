package middleware

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Middleware func(context.Context, *mongo.Database) context.Context

type MiddlewareKey string
