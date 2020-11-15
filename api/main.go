package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dominik-zeglen/geralt/core"
	"github.com/dominik-zeglen/geralt/models"
	"github.com/dominik-zeglen/geralt/utils"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	secret string
	port   string
}

type API struct {
	conf   config
	db     *mongo.Database
	geralt core.Core
	users  *cache.Cache
}

func (api *API) Init() {
	api.conf.port = utils.GetEnvOrPanic("PORT")
	api.conf.secret = utils.GetEnvOrPanic("SECRET")
	dbConfig := models.GetDBConfig()

	client, err := mongo.
		NewClient(
			options.
				Client().
				ApplyURI(dbConfig.Hostname).
				SetAuth(options.Credential{
					Username: dbConfig.Username,
					Password: dbConfig.Password,
				}))

	if err != nil {
		panic(err)
	}

	dbCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(dbCtx)

	if err != nil {
		panic(err)
	}

	api.db = client.Database(dbConfig.DBName)

	api.users = cache.New(2*time.Minute, 4*time.Minute)

	api.geralt = core.Core{}
	api.geralt.Init(api.db)
}

func (api API) Start() {
	mux := http.NewServeMux()
	mux.Handle("/", combineMiddlewares(
		http.HandlerFunc(api.handleReply),
		[]Middleware{
			api.withTracing,
			api.withJwt,
			api.withUser,
			api.withBot,
		},
	),
	)
	mux.HandleFunc("/auth", api.withTracing(api.handleAuth))

	log.Println("Serving at port", api.conf.port)
	err := http.ListenAndServe(":"+api.conf.port, mux)
	log.Fatal(err)
}
