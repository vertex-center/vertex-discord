package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/quentinguidee/microservice-core/pubsub"
	"github.com/quentinguidee/microservice-core/router"
)

func initializeRouter() *gin.Engine {
	r := router.CreateRouter()

	subscriber := pubsub.Sub("SPOTIFY_PLAYER_CHANGE")

	go func() {
		for {
			msg, err := subscriber.ReceiveMessage(context.Background())
			if err != nil {
				panic(err)
			}

			println(msg.String())
		}
	}()

	return r
}
