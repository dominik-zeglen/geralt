package api

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (api *API) withUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(handlers.UserContextKey).(models.User)

		collection := api.db.Collection(models.UsersCollectionKey)
		err := collection.FindOne(context.TODO(), bson.M{
			"_id": user.ID,
		}).Decode(&user)

		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), handlers.UserContextKey, user)

		next(w, r.WithContext(ctx))
	}
}
