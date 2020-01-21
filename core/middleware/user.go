package middleware

import (
	"context"

	"github.com/dominik-zeglen/geralt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserContextKey = MiddlewareKey("user")

func WithUser(ctx context.Context, db *mongo.Database) context.Context {
	user := ctx.Value(UserContextKey).(models.User)

	collection := db.Collection(models.UsersCollectionKey)
	err := collection.FindOne(context.TODO(), bson.M{
		"_id": user.ID,
	}).Decode(&user)

	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, UserContextKey, user)
}
