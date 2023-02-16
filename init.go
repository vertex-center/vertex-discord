package main

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/quentinguidee/microservice-core/pubsub"
	"log"
)

var environment Environment

type Environment struct {
	DiscordToken     string `env:"DISCORD_TOKEN,required"`
	SpotifyEmojiID   string `env:"SPOTIFY_EMOJI_ID,required"`
	SpotifyEmojiName string `env:"SPOTIFY_EMOJI_ID,required"`
}

func main() {
	loadEnv()

	pubsub.InitPubSub()

	r := initializeRouter()

	err := r.Run(":6151")
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = env.Parse(&environment)
	if err != nil {
		log.Fatalf("Failed to parse .env to Environment: %v", err)
	}
}
