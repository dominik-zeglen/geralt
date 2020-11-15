package api

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addUserToContext(
	ctx context.Context,
	user *handlers.User,
) context.Context {
	return context.WithValue(ctx, handlers.UserContextKey, user)
}

func (api *API) withUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		span, _ := opentracing.StartSpanFromContext(
			r.Context(),
			"middleware-user",
		)

		userID, ok := ctx.Value(userIDContextKey).(string)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
		}

		cachedUser := api.getUser(userID)
		if cachedUser != nil {
			span.LogFields(
				log.Bool("cached", true),
			)
			ctx = addUserToContext(ctx, cachedUser)

			span.Finish()
			next(w, r.WithContext(ctx))
			return
		} else {
			id, conversionErr := primitive.ObjectIDFromHex(userID)
			if conversionErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
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

			span.LogFields(
				log.Bool("cached", false),
			)
			ctx = addUserToContext(ctx, rememberedUser)

			span.Finish()
			next(w, r.WithContext(ctx))
			return
		}
	}
}
