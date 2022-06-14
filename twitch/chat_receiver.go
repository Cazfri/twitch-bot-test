package twitch

import (
	"log"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
)

type ChatReceiver struct {
	twitchClient    *twitch.Client
	messages        chan string
	allowedCommands map[string]bool
}

func NewChatReceiver(chatToWatch string, allowedCommands map[string]bool, messagesChan chan string) *ChatReceiver {
	twitchClient := twitch.NewAnonymousClient()

	receiver := ChatReceiver{
		twitchClient:    twitchClient,
		messages:        messagesChan,
		allowedCommands: allowedCommands,
	}

	twitchClient.OnPrivateMessage(receiver.handleMessage)

	twitchClient.Join(chatToWatch)

	return &receiver
}

func (r *ChatReceiver) Connect() error {
	return r.twitchClient.Connect()
}

func (r *ChatReceiver) handleMessage(twitchMessage twitch.PrivateMessage) {
	message := twitchMessage.Message
	log.Println("Twitch received message", message)

	if !r.commandIsAllowed(message) {
		return
	}
	// TODO: Right now a filled buffer will block, but should it just fail?
	r.messages <- message
}

func (r *ChatReceiver) commandIsAllowed(command string) bool {
	cleanCommand := cleanCommand(command)

	_, allowed := r.allowedCommands[cleanCommand]
	return allowed
}

func cleanCommand(command string) string {
	command = strings.ToLower(command)
	return command
}
