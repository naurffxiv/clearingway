package discord

import "github.com/bwmarrin/discordgo"

type Discord struct {
	Session *discordgo.Session
}

// NewDiscord - Initializes a new Discord session with the provided bot token
func NewDiscord(token string) (*Discord, error) {
	// "Bot " prefix is required for bot authentication
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Discord{Session: session}, nil
}
