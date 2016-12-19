package main

import (
	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/Config"
	"github.com/philbrookes/game-server/HttpController"
	"gopkg.in/mgo.v2"
)

func setupHTTPServer(session *mgo.Session, cfg *Config.Config) *mux.Router {

	r := mux.NewRouter()

	// user routes
	r.HandleFunc("/user", HttpController.NewUser(session, cfg)).Methods("GET")
	r.HandleFunc("/user/{id}", HttpController.NewUser(session, cfg)).Methods("GET", "PUT", "POST", "DELETE")

	return r
}
