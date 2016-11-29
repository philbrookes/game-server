package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	http.Handle("/connect", websocket.Handler(gameServer))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/user", getUserHandler(session))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendJSON(rw http.ResponseWriter, res interface{}) error {
	textRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	headers := rw.Header()
	headers.Add("Content-Type", "Application/json")
	_, err = rw.Write(textRes)
	return err
}
