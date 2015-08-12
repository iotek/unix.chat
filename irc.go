package main

import (
	"fmt"
	"net/http"
	"hash/fnv"
	"golang.org/x/net/websocket"
)

var (
	clients = make(map[ClientConn]int)
	message = websocket.Message
)

type ClientConn struct {
	websocket *websocket.Conn
	ip  string
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())[0:5]
}

func ircHandler(ws *websocket.Conn) {
	var err error
	var clientmessage string

	current := ClientConn{ws, ws.Request().RemoteAddr}
	clients[current] = 0

	fmt.Println("number of clients connected ", len(clients))

	for {
		if err = message.Receive(ws, &clientmessage); err != nil {
			delete(clients, current)
			fmt.Println("client disconnect,", len(clients), " clients remain")
			return
		}

		clientmessage = "<" + hash(current.ip) + "> " + clientmessage

		for cs, _ := range clients {
			if err = message.Send(cs.websocket, clientmessage); err != nil {
				 fmt.Println("sending failed")
			}
		}
	}
}

func main() {
	http.Handle("/irc", websocket.Handler(ircHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

