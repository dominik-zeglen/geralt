package api

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
)

func (api *API) withBot(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bot, err := models.GetBot(api.db)
		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), handlers.BotContextKey, bot)

		next(w, r.WithContext(ctx))
	}
}
