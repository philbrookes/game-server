package httpController

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	mgo "gopkg.in/mgo.v2"

	"strconv"

	"github.com/philbrookes/game-server/config"
	"github.com/philbrookes/game-server/user"
)

// NewUser returns a new user controller for handling requests related to single user
func NewUser(session *mgo.Session, cfg *config.Config) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		sender, err := GetSender(cfg.OutputFormat)
		if err != nil {
			log.Fatal(err)
		}
		switch r.Method {
		case "GET":
			getUser(rw, r, session, cfg, sender)
		case "POST":
			createUser(rw, r, session, cfg, sender)
		case "PUT":
			updateUser(rw, r, session, cfg, sender)
		case "DELETE":
			deleteUser(rw, r, session, cfg, sender)
		}
	}
}

func getUser(rw http.ResponseWriter, r *http.Request, session *mgo.Session, cfg *config.Config, sender Sender) {
	log.Print("getting users")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	payload := user.User{
		ID: id,
	}
	users, err := user.GetUsers(payload, session, cfg)
	if err != nil {
		log.Println(err)
		sender(rw, err.Error())
		return
	}

	sender(rw, user.PublicFilter(users))
}

func updateUser(rw http.ResponseWriter, r *http.Request, session *mgo.Session, cfg *config.Config, sender Sender) {
	log.Println("Updating users")
}

func createUser(rw http.ResponseWriter, r *http.Request, session *mgo.Session, cfg *config.Config, sender Sender) {
	log.Println("Creating users")

	decoder := json.NewDecoder(r.Body)
	newUser := user.User{}
	if err := decoder.Decode(&newUser); err != nil {
		log.Fatalln(err)
	}

	user, err := user.CreateUser(newUser, session, cfg)

	if err != nil {
		sender(rw, err.Error())
		log.Println(err)
		return
	}

	sender(rw, user)
}

func deleteUser(rw http.ResponseWriter, r *http.Request, session *mgo.Session, cfg *config.Config, sender Sender) {
	log.Println("Deleting users")
}

func getUsers(rw http.ResponseWriter, r *http.Request, userCollection *mgo.Collection) {
	log.Println("Getting users")
	users := user.Users{}
	userCollection.Find(nil).All(&users)
	if err := sendJSON(rw, users); err != nil {
		log.Fatalln(err)
	}
}
