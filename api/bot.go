package api

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
	opentracing "github.com/opentracing/opentracing-go"
)

func (api *API) withBot(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, ctx := opentracing.StartSpanFromContext(
			r.Context(),
			"middleware-bot",
		)

		bot, err := models.GetBot(api.db)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx = context.WithValue(ctx, handlers.BotContextKey, bot)

		span.Finish()
		next(w, r.WithContext(ctx))
	}
}
