package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	FFLOGS_CLIENT_ID     string
	FFLOGS_CLIENT_SECRET string
	DISCORD_TOKEN        string
	ENV                  string
}

func LoadEnv() *Env {
	// Check if we're in production mode
	env := os.Getenv("ENV")
	if env != "production" {
		// In development, load from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Could not load .env file: %v", err)
		}
	}

	// -------------- LOAD ENV VARIABLES --------------
	fflogsClientId, ok := os.LookupEnv("FFLOGS_CLIENT_ID")
	if !ok {
		log.Fatalf("FFLOGS_CLIENT_ID is not set in the environment")
	}
	fflogsClientSecret, ok := os.LookupEnv("FFLOGS_CLIENT_SECRET")
	if !ok {
		log.Fatalf("FFLOGS_CLIENT_SECRET is not set in the environment")
	}
	discordToken, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		log.Fatalf("DISCORD_TOKEN is not set in the environment")
	}

	// -------------- RETURN ENV --------------
	return &Env{
		FFLOGS_CLIENT_ID:     fflogsClientId,
		FFLOGS_CLIENT_SECRET: fflogsClientSecret,
		DISCORD_TOKEN:        discordToken,
		ENV:                  env,
	}
}
