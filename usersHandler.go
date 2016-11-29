package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"encoding/json"

	mgo "gopkg.in/mgo.v2"
)

//User definition in Mongo
type User struct {
	ID        int    `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func getUserHandler(session *mgo.Session) http.HandlerFunc {
	userCollection := session.DB("game").C("user")
	return func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUsers(rw, r, userCollection)
		case "POST":
			createUser(rw, r, userCollection)
		case "PUT":
			updateUser(rw, r, userCollection)
		case "DELETE":
			deleteUser(rw, r, userCollection)
		}
	}

}

func getUsers(rw http.ResponseWriter, r *http.Request, userCollection *mgo.Collection) {
	log.Println("Getting users")
	users := []User{}
	userCollection.Find(nil).All(&users)
	if err := sendJSON(rw, users); err != nil {
		log.Fatalln(err)
	}
}

func updateUser(rw http.ResponseWriter, r *http.Request, userCollection *mgo.Collection) {
	log.Println("Updating users")
}

func createUser(rw http.ResponseWriter, r *http.Request, userCollection *mgo.Collection) {
	log.Println("Creating users")

	decoder := json.NewDecoder(r.Body)
	user := User{}
	if err := decoder.Decode(&user); err != nil {
		log.Fatalln(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	user.Password = string(hashedPassword)

	if err := userCollection.Insert(&user); err != nil {
		log.Fatalln(err)
	}

	if err := sendJSON(rw, user); err != nil {
		log.Fatalln(err)
	}
}

func deleteUser(rw http.ResponseWriter, r *http.Request, userCollection *mgo.Collection) {
	log.Println("Deleting users")
}
