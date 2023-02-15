package main

import "log"

func main() {
	r := initializeRouter()

	err := r.Run(":6151")
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
