package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/Config"
	"github.com/philbrookes/game-server/HttpController"
	"gopkg.in/mgo.v2"
)

func setupHTTPServer(r *mux.Router, session *mgo.Session, cfg *Config.Config) *mux.Router {
	// user routes
	r.HandleFunc("/user", HttpController.NewUser(session, cfg)).Methods("GET")
	r.HandleFunc("/user/{id}", HttpController.NewUser(session, cfg)).Methods("GET", "PUT", "POST", "DELETE")

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	return r
}
