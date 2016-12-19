package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/Config"
	"github.com/philbrookes/game-server/HttpController"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2"
)

func gameServer(ws *websocket.Conn) {
	for {
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Receive: %s\n", msg[:n])

		m, err := ws.Write(msg[:n])
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Send: %s\n", msg[:m])
	}
}

func main() {
	cfg, err := Config.Get(os.Getenv("GAME_ENVIRONMENT"))
	if err != nil {
		panic(err)
	}
	session, err := mgo.Dial(cfg.MongoHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	r := mux.NewRouter()
	r.Handle("/connect", websocket.Handler(gameServer))

	// user routes
	r.HandleFunc("/user", HttpController.NewUser(session, cfg)).Methods("GET")
	r.HandleFunc("/user/{id}", HttpController.NewUser(session, cfg)).Methods("GET", "PUT", "POST", "DELETE")

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/public"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}
