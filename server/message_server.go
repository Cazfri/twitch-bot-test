package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type MessageServer struct {
	messageChan chan string
	port        int

	upgrader websocket.Upgrader
}

type Message struct {
	Message string
}

func NewMessageServer(messageChan chan string, port int) *MessageServer {
	server := MessageServer{
		messageChan: messageChan,
		port:        port,
		upgrader:    websocket.Upgrader{},
	}

	return &server
}

func (s *MessageServer) Serve() error {
	http.HandleFunc("/messages", s.handleMessages)
	return http.ListenAndServe(fmt.Sprintf("localhost:%d", s.port), nil)
}

func (s *MessageServer) handleMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("Client connected, upgrading to ws connection")
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer c.Close()

	log.Println("Sending messages")
	for {
		message := <-s.messageChan
		log.Println("Message server sending message:", message)
		if err := c.WriteJSON(Message{Message: message}); err != nil {
			log.Println("Error writing JSON:", err)
			// TODO: This will break out when client is disconnected, but I need a better way to detect this
			break
		}
	}
}
