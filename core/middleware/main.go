package middleware

import "context"

type Middleware func(context.Context) context.Context

type MiddlewareKey string
