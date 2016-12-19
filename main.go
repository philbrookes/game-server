package main

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/Config"
	"gopkg.in/mgo.v2"
)

func main() {
	cfg, err := Config.Get(os.Getenv("GAME_ENVIRONMENT"))
	if err != nil {
		panic(err)
	}

	session := connectToMongo(cfg)
	defer session.Close()

	r := mux.NewRouter()

	setupSocketServer(r, session, cfg)
	setupHTTPServer(r, session, cfg)

	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, r))
}

func connectToMongo(cfg *Config.Config) *mgo.Session {
	session, err := mgo.Dial(cfg.MongoHost)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}
