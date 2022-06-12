package main

import (
	"log"

	"github.com/cazfri/twitch-bot-test/server"
	"github.com/cazfri/twitch-bot-test/twitch"
)

func main() {
	// TODO: Currently connections will receive subsets of the message stream, which may not be correct.
	// When clients connect, they should have to register with the chat receiver, which will create them a new
	// channel to consume from
	// What should probably happen is that the chat receiver creates a channel in a map keyed off of
	// a key given in the register call (or doesn't if it already exists, this key acts as a sort of
	// consumer group) (consumer group can also be chosen by a url parameter by the connecting client),
	// then returns a handle object that holds the key and provides access to the channel. Client
	// then must defer handle.Close() to close out the handle, which closes the channel and removes
	// it from the map.
	messageBufferSize := 100
	messages := make(chan string, messageBufferSize)

	chatReceiver := twitch.NewChatReceiver(messages)
	go func() {
		log.Println("Connecting to twitch client")
		if err := chatReceiver.Connect(); err != nil {
			log.Fatal("Cannot connect to twitch client", err)
		}
	}()

	port := 8080
	log.Println("Starting message server on port ", port)
	messageServer := server.NewMessageServer(messages, port)
	if err := messageServer.Serve(); err != nil {
		log.Fatal("Cannot serve message server", err)
	}
}
