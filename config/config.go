package config

// TODO: All of this should be inside a JSON, TOML, or YAML file and loaded when the service starts,
// but this is a quick and dirty way to separate the pieces first.

const (
	MessageServerPort = 8080
	TwitchChatToWatch = "cazfri"
)

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
