package twitch

import (
	"log"

	"github.com/gempir/go-twitch-irc/v3"
)

type ChatReceiver struct {
	twitchClient *twitch.Client
	messages     chan string
}

func NewChatReceiver(messagesChan chan string) *ChatReceiver {

	twitchClient := twitch.NewAnonymousClient()

	receiver := ChatReceiver{
		twitchClient: twitchClient,
		messages:     messagesChan,
	}

	twitchClient.OnPrivateMessage(receiver.handleMessage)

	twitchClient.Join("cazfri")

	return &receiver
}

func (r *ChatReceiver) Connect() error {
	return r.twitchClient.Connect()
}

func (r *ChatReceiver) handleMessage(twitchMessage twitch.PrivateMessage) {
	message := twitchMessage.Message
	log.Println("Twitch received message", message)

	// TODO: Right now a filled buffer will block, but should it just fail?
	r.messages <- message
}
