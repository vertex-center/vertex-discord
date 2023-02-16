package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quentinguidee/microservice-core/pubsub"
	"github.com/quentinguidee/microservice-core/router"
	"net/http"
	"strings"
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

			body := fmt.Sprintf(`{"custom_status":{"emoji_name":"spotify","emoji_id":"%s"}}`, environment.SpotifyEmojiID)

			token := fmt.Sprintf("%s", environment.DiscordToken)

			req, err := http.NewRequest("PATCH", "https://discord.com/api/users/@me/settings", strings.NewReader(body))
			req.Header.Set("Authorization", token)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "*/*")

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			println(token)
			println(res.Status)

			println(msg.String())
		}
	}()

	return r
}
