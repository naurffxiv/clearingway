package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	CONFIG_PATH          string
	FFLOGS_CLIENT_ID     string
	FFLOGS_CLIENT_SECRET string
	DISCORD_TOKEN        string
	ENV                  string
}

// LoadEnv - Loads environment variables from either the system or a .env file
func LoadEnv() (*Env, error) {
	// Check if we're in production mode
	env := os.Getenv("ENV")
	if env != "production" {
		// In development, load from .env file
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	// -------------- LOAD ENV VARIABLES --------------
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		return nil, errors.New("CONFIG_PATH not set in environment")
	}
	fflogsClientId, ok := os.LookupEnv("FFLOGS_CLIENT_ID")
	if !ok {
		return nil, errors.New("FFLOGS_CLIENT_ID not set in environment")
	}
	fflogsClientSecret, ok := os.LookupEnv("FFLOGS_CLIENT_SECRET")
	if !ok {
		return nil, errors.New("FFLOGS_CLIENT_SECRET not set in environment")
	}
	discordToken, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		return nil, errors.New("DISCORD_TOKEN not set in environment")
	}

	// -------------- RETURN ENV --------------
	return &Env{
		CONFIG_PATH:          configPath,
		FFLOGS_CLIENT_ID:     fflogsClientId,
		FFLOGS_CLIENT_SECRET: fflogsClientSecret,
		DISCORD_TOKEN:        discordToken,
		ENV:                  env,
	}, nil
}
