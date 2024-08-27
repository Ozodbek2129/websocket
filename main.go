package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler_conection)
	fmt.Println("Server Started >>> localhost:3000")
	err := http.ListenAndServe("18.171.153.211:3000", mux)
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func Handler_conection(w http.ResponseWriter, r *http.Request) {
	websocketCon, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error: ", err)
	}

	address := websocketCon.NetConn()
	defer websocketCon.Close()

	for {
		var msg Message
		err := websocketCon.ReadJSON(&msg)
		if err != nil {
			log.Fatal("error: ", err)
		}

		log.Println("Message from user: ", msg.Name, "Address: ", address.RemoteAddr(), "Message: ", msg.Message)

		err = websocketCon.WriteJSON(&msg)
		if err != nil {
			log.Fatal("error: ", err)
		}
	}
}
