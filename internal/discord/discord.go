package discord

import "github.com/bwmarrin/discordgo"

type Discord struct {
	Session *discordgo.Session
}

func NewDiscord(token string) (*Discord, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Discord{Session: session}, nil
}
