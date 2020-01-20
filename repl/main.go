package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dominik-zeglen/geralt/core"
	"github.com/dominik-zeglen/geralt/core/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repl struct {
	db *mongo.Database
}

func (r *Repl) Init() {
	client, err := mongo.
		NewClient(
			options.
				Client().
				ApplyURI("mongodb://localhost:27017").
				SetAuth(options.Credential{
					Username: "geralt",
					Password: "geralt",
				}))

	if err != nil {
		panic(err)
	}

	dbCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(dbCtx)

	if err != nil {
		panic(err)
	}

	db := client.Database("geralt")

	r.db = db
}

func getCurrentUser(users []middleware.UserData, reader *bufio.Reader) primitive.ObjectID {
	for userIndex, user := range users {
		fmt.Printf("%d) %s\n", userIndex+1, user.Name)
	}

	var userID primitive.ObjectID

	fmt.Println(users)

	for ok := false; !ok; {
		userIndexInput, _ := reader.ReadString('\n')
		userIndex, convErr := strconv.Atoi(strings.Trim(userIndexInput, "\n"))
		userID = users[userIndex-1].ID

		ok = convErr == nil
	}

	return userID
}

func (r Repl) Start() {
	geralt := core.Core{}
	geralt.Init()

	reader := bufio.NewReader(os.Stdin)

	middlewares := []middleware.Middleware{
		middleware.WithBot,
		middleware.WithUser,
	}

	var userID primitive.ObjectID
	users := []middleware.UserData{}
	maxUsers := int64(20)
	userCursor, fetchErr := r.
		db.
		Collection("users").
		Find(
			context.TODO(),
			bson.D{},
			&options.FindOptions{
				Limit: &maxUsers,
			})
	if fetchErr != nil {
		panic(fetchErr)
	}

	decodeErr := userCursor.All(context.TODO(), &users)
	if decodeErr != nil {
		panic(decodeErr)
	}

	if len(users) > 0 {
		userID = getCurrentUser(users, reader)
	} else {
		user, insertError := r.
			db.
			Collection("users").
			InsertOne(
				context.TODO(),
				middleware.UserData{
					Name:  "Admin",
					Email: "admin@example.com",
				})

		if insertError != nil {
			panic(insertError)
		}

		userID = user.InsertedID.(primitive.ObjectID)
	}

	for true {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		ctx := context.WithValue(
			context.Background(),
			middleware.UserContextKey,
			middleware.UserData{
				ID: userID,
			})
		for _, withMiddleware := range middlewares {
			ctx = withMiddleware(ctx, r.db)
		}

		fmt.Printf("%s\n", geralt.Reply(ctx, text))
	}
}
