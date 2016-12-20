package socketController

import (
	"log"
	"strconv"

	mgo "gopkg.in/mgo.v2"

	"github.com/philbrookes/game-server/config"
	"github.com/philbrookes/game-server/user"
	"golang.org/x/net/websocket"
)

//UserList sends a list of all users to the provided websocket
func UserList(ws *websocket.Conn, cmd Command, session *mgo.Session, cfg *config.Config, sender Sender) {
	log.Print("getting users")

	id, _ := strconv.Atoi(cmd[1])
	payload := user.User{
		ID: id,
	}
	users, err := user.GetUsers(payload, session, cfg)
	if err != nil {
		log.Println(err)
		sender(ws, "ERROR: "+err.Error())
		return
	}

	sender(ws, user.PublicFilter(users))
}
