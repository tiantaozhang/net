package main

import (
	//"github.com/golang/net/websocket"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	for {
		fmt.Println("websocket...")
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("receive back from client:" + reply)
		msg := "received: " + reply
		fmt.Println("sending to client: " + msg)
		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func main() {
	http.Handle("/websocket", websocket.Handler(Echo))
	http.Handle("/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(":4321", nil); err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
