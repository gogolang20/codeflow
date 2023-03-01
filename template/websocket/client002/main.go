package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	go send(conn)

	for {
		message, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(message, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
