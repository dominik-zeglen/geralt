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
	"github.com/dominik-zeglen/geralt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getEnvOrPanic(key string) string {
	r := os.Getenv(key)
	if r == "" {
		panic(fmt.Errorf("Environment variable %s not set", key))
	}

	return r
}

type dbConfig struct {
	username string
	password string
	hostname string
	dbname   string
}

type config struct {
	db dbConfig
}

func getConfig() config {
	return config{
		db: dbConfig{
			username: getEnvOrPanic("DB_USERNAME"),
			password: getEnvOrPanic("DB_PASSWORD"),
			hostname: getEnvOrPanic("DB_HOSTNAME"),
			dbname:   getEnvOrPanic("DB_NAME"),
		},
	}
}

type Repl struct {
	db *mongo.Database
}

func (r *Repl) Init() {
	conf := getConfig()

	client, err := mongo.
		NewClient(
			options.
				Client().
				ApplyURI(conf.db.hostname).
				SetAuth(options.Credential{
					Username: conf.db.username,
					Password: conf.db.password,
				}))

	if err != nil {
		panic(err)
	}

	dbCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(dbCtx)

	if err != nil {
		panic(err)
	}

	db := client.Database(conf.db.dbname)

	r.db = db
}

func getCurrentUser(users []models.User, reader *bufio.Reader) primitive.ObjectID {
	fmt.Println("Select user:")
	for userIndex, user := range users {
		fmt.Printf("%d) %s\n", userIndex+1, user.Name)
	}

	var userID primitive.ObjectID

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

	users := []models.User{}
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

	userID := getCurrentUser(users, reader)

	for true {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		user := models.User{}
		user.ID = userID

		ctx := context.WithValue(
			context.Background(),
			middleware.UserContextKey,
			user,
		)
		for _, withMiddleware := range middlewares {
			ctx = withMiddleware(ctx, r.db)
		}

		fmt.Printf("%s\n", geralt.Reply(ctx, text))
	}
}
