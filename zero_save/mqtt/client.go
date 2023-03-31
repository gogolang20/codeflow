package main

import (
	"fmt"
	"net"
	"time"
)

type MqttServer struct {
	port       int
	clients    []*MqttClient
	newClients chan net.Conn
}

type MqttClient struct {
	conn   net.Conn
	reader *MqttReader
}

func NewMqttReader(conn net.Conn) *MqttReader {
	return &MqttReader{}
}

type MqttReader struct {
}

func (m *MqttReader) Read() (string, error) {

}

func NewMqttServer(port int) *MqttServer {
	return &MqttServer{port: port, clients: make([]*MqttClient, 0), newClients: make(chan net.Conn)}
}

func (s *MqttServer) Start() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%!d(MISSING)", s.port))
	if err != nil {
		fmt.Printf("failed to start server: %!s(MISSING)\n", err.Error())
		return
	}
	fmt.Printf("server started on port %!d(MISSING)\n", s.port)

	go s.acceptNewClients(ln)

	for {
		for _, client := range s.clients {
			message, err := client.reader.Read()
			if err != nil {
				fmt.Printf("failed to read message: %!s(MISSING)\n", err.Error())
			} else {
				fmt.Printf("received message: %!s(MISSING)\n", message)
			}
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (s *MqttServer) acceptNewClients(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("failed to accept new client: %!s(MISSING)\n", err.Error())
		} else {
			fmt.Printf("new client connected: %!s(MISSING)\n", conn.RemoteAddr().String())
			client := &MqttClient{
				conn:   conn,
				reader: NewMqttReader(conn),
			}
			s.clients = append(s.clients, client)
			s.newClients <- conn
		}
	}
}

func main() {
	server := NewMqttServer(1883)

	server.Start()
}
