package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Message struct {
	Msg    []byte
	Sender *Client
}

type Client struct {
	Socket net.Conn
	Name   string
	Id     int
}

func main() {
	clients := []Client{}
	listener, err := net.Listen("tcp", ":27072")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	messages := make(chan Message)

	go func() {
		for message := range messages {
			for _, client := range clients {
				if client.Id != message.Sender.Id {
					client.Socket.Write([]byte(message.Sender.Name + " <-- : " + string(message.Msg)))
					fmt.Print(strconv.Itoa(message.Sender.Id) + " <-- : " + string(message.Msg))
				}
			}
		}
	}()

	for {
		newConn, err := listener.Accept()
		newClient := Client{Socket: newConn, Id: len(clients), Name: "player_" + strconv.Itoa(len(clients))}
		clients = append(clients, newClient)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		go func(client Client) {
			for {
				msg, err := bufio.NewReader(client.Socket).ReadBytes('\n')
				if err != nil {
					fmt.Fprint(os.Stderr, err.Error())
					continue
				}
				fmt.Print(strconv.Itoa(client.Id) + " --> : " + string(msg))
				messages <- Message{Msg: msg, Sender: &client}
			}
		}(newClient)
	}
}
