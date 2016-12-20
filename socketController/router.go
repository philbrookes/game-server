package socketController

import (
	"strings"

	"github.com/philbrookes/game-server/config"
	"golang.org/x/net/websocket"
	mgo "gopkg.in/mgo.v2"
)

//Routes defines a map of command strings to controller functions
type Routes map[string]Controller

//Router centralises calls to add and remove command controllers
type Router struct {
	routes Routes
}

//NewRouter creates and returns a pointer to a new router
func NewRouter() *Router {
	router := Router{routes: Routes{}}
	return &router
}

//Command is a slice containing the initial command plus it's arguments
type Command []string

//Controller is a function which can deal with a specific command
type Controller func(*websocket.Conn, Command, *mgo.Session, *config.Config, Sender)

//AddRoute the route with the controller
func (r *Router) AddRoute(route string, controller Controller) {
	r.routes[route] = controller
}

//GetController will find and return the controller for the given command, or unknownCommand
func (r *Router) GetController(route string) Controller {
	if val, ok := r.routes[route]; ok {
		return val
	}

	return unknownCommand
}

func unknownCommand(ws *websocket.Conn, cmd Command, session *mgo.Session, cfg *config.Config, sender Sender) {
	sender(ws, "Unknown command: "+strings.Join(cmd, " "))
}
