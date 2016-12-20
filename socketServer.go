package main

import (
	"log"
	"strings"

	"github.com/gorilla/mux"
	"github.com/philbrookes/game-server/config"
	"github.com/philbrookes/game-server/socketController"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2"
)

func setupSocketServer(r *mux.Router, session *mgo.Session, cfg *config.Config) *mux.Router {

	socketRouter := socketController.NewRouter()
	socketRouter.AddRoute("user_list", socketController.UserList)

	r.Handle("/"+cfg.WebsocketListener, websocket.Handler(gameServer(session, cfg, socketRouter)))

	return r
}

func gameServer(session *mgo.Session, cfg *config.Config, sr *socketController.Router) func(*websocket.Conn) {
	return func(ws *websocket.Conn) {
		for {
			msg := make([]byte, 512)
			n, err := ws.Read(msg)
			if err != nil {
				log.Println(err)
				return
			}
			cmd := string(msg[:n])
			cmdParts := strings.Fields(cmd)

			controller := sr.GetController(cmdParts[0])

			sender, err := socketController.GetSender(cfg.OutputFormat)
			if err != nil {
				panic(err)
			}

			controller(ws, cmdParts, session, cfg, sender)
		}
	}
}
