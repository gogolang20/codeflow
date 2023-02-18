package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conns = append(conns, conn)

	for {
		message, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("Are you say "+string(p)+" ?"))
		}
		log.Println(message, string(p))
	}
	defer conn.Close()
	log.Println("service close...")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
