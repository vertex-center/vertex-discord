package main

import (
	"context"
	"encoding/json"
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

			println(msg.String())

			var messageJSON map[string]interface{}
			err = json.Unmarshal([]byte(msg.Payload), &messageJSON)
			if err != nil {
				fmt.Printf("Failed to parse message to JSON: %v\n", err)
				continue
			}

			var body string
			if messageJSON["is_playing"] == true {
				//artist := messageJSON["track"].(map[string]interface{})["artist"]
				body = fmt.Sprintf(
					`{"custom_status":{"emoji_name":"%s","emoji_id":"%s","text":%s}}`,
					environment.SpotifyEmojiName,
					environment.SpotifyEmojiID,
					//fmt.Sprintf("\"Écoute %s\"", artist),
					"null",
				)
			} else {
				body = `{"custom_status":{"emoji_name":null,"emoji_id":null,"text":null}}`
			}

			token := fmt.Sprintf("%s", environment.DiscordToken)

			req, err := http.NewRequest("PATCH", "https://discord.com/api/users/@me/settings", strings.NewReader(body))
			req.Header.Set("Authorization", token)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "*/*")

			client := &http.Client{}
			_, err = client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()

	return r
}
