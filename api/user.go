package api

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *API) withUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(userIDContextKey).(string)
		if !ok {
			// return 400
		}

		cachedUser := api.getUser(userID)
		if cachedUser != nil {
			ctx := context.WithValue(
				r.Context(),
				handlers.UserContextKey,
				*cachedUser,
			)

			next(w, r.WithContext(ctx))
		} else {
			id, conversionErr := primitive.ObjectIDFromHex(userID)
			if conversionErr != nil {
				// return 500
			}

			var user models.User
			collection := api.db.Collection(models.UsersCollectionKey)
			selectionErr := collection.FindOne(context.TODO(), bson.M{
				"_id": id,
			}).Decode(&user)

			if selectionErr != nil {
				panic(selectionErr)
			}

			rememberedUser := api.rememberUser(user)

			ctx := context.WithValue(
				r.Context(),
				handlers.UserContextKey,
				rememberedUser,
			)

			next(w, r.WithContext(ctx))
		}
	}
}
