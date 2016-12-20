package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/config"
	"github.com/philbrookes/game-server/httpController"
	mgo "gopkg.in/mgo.v2"
)

func setupHTTPServer(r *mux.Router, session *mgo.Session, cfg *config.Config) *mux.Router {
	// user routes
	r.HandleFunc("/user", httpController.NewUser(session, cfg)).Methods("GET")
	r.HandleFunc("/user/{id}", httpController.NewUser(session, cfg)).Methods("GET", "PUT", "POST", "DELETE")

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	return r
}
