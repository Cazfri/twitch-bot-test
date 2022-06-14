package config

const MessageServerPort = 8080

func AllowedCommands() map[string]bool {
	return map[string]bool{
		"up":       true,
		"down":     true,
		"forward":  true,
		"backward": true,
		"left":     true,
		"right":    true,
	}
}
