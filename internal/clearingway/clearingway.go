package clearingway

import (
	"clearingway/internal/clearingway/config"
	"clearingway/internal/discord"
	"clearingway/internal/env"
)

type Clearingway struct {
	Config  *config.BotConfig
	Discord *discord.Discord
}

func NewClearingway(env *env.Env) (*Clearingway, error) {
	loadedConfig, err := config.Init(env.CONFIG_PATH)
	if err != nil {
		return nil, err
	}

	discordClient, err := discord.NewDiscord(env.DISCORD_TOKEN)
	if err != nil {
		return nil, err
	}

	return &Clearingway{
		Config:  loadedConfig,
		Discord: discordClient,
	}, nil
}

func (cw *Clearingway) Start() error {
	return cw.Discord.Session.Open()
}

func (cw *Clearingway) Stop() error {
	return cw.Discord.Session.Close()
}

func (cw *Clearingway) GetConfig() *config.BotConfig {
	return cw.Config
}

func (cw *Clearingway) GetDiscord() *discord.Discord {
	return cw.Discord
}
