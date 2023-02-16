package main

import (
	"github.com/quentinguidee/microservice-core/pubsub"
	"log"
)

func main() {
	pubsub.InitPubSub()

	r := initializeRouter()

	err := r.Run(":6151")
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
