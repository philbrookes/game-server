package User

import (
	"log"

	"github.com/philbrookes/game-server/Config"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
)

//User definition in Mongo
type User struct {
	ID        int    `json:"id,omitempty" bson:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Username  string `json:"username,omitempty" bson:"username,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
}

//PublicFilter removes any fields that should not be displayed publically
func PublicFilter(users Users) Users {
	filteredUsers := Users{}

	for _, user := range users {
		user.Password = "*** ***"
		filteredUsers = append(filteredUsers, user)
	}

	return filteredUsers
}

//Users is a slice of User
type Users []User

//GetUsers looks in the users table in the session for users as defined in the payload
func GetUsers(payload User, session *mgo.Session, cfg *Config.Config) (Users, error) {
	users := Users{}
	err := session.DB(cfg.DatabaseName).C(cfg.UserTable).Find(payload).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

//CreateUser will create a user defined in the payload in the users table in the provided session
func CreateUser(user User, session *mgo.Session, cfg *Config.Config) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	user.Password = string(hashedPassword)

	if err := session.DB(cfg.DatabaseName).C(cfg.UserTable).Insert(&user); err != nil {
		log.Fatalln(err)
	}

	return user, nil
}
