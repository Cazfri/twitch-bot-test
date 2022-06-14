package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// MessageServer is the singleton server that serves messages to
type MessageServer struct {
	messageChan chan string
	port        int

	upgrader websocket.Upgrader
}

// Message defines the structure of the messages sent to clients
type Message struct {
	Message string
}

// NewMessageServer creates a new message server, but does not start it
func NewMessageServer(messageChan chan string, port int) *MessageServer {
	server := MessageServer{
		messageChan: messageChan,
		port:        port,
		upgrader:    websocket.Upgrader{},
	}

	return &server
}

// Serve starts the message server, and will block unless the server encounters a fatal error
func (s *MessageServer) Serve() error {
	log.Println("Starting message server on port", s.port)
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
			// TODO: This will break out when client is disconnected, but there is probably a better
			// way to detect a disconnected client
			break
		}
	}
}
