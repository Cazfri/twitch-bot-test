package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/messages"
	log.Println("Dialling ", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dail error:", err)
	}
	defer c.Close()

	log.Println("Waiting for messages...")
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
		}
		log.Println("recv:", string(message))
	}
}
