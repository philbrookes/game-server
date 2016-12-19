package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/Config"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2"
)

func setupSocketServer(r *mux.Router, session *mgo.Session, cfg *Config.Config) *mux.Router {

	r.Handle("/"+cfg.WebsocketListener, websocket.Handler(gameServer))

	return r
}

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
