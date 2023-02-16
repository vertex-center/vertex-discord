# Discord Service `BETA`

A Discord microservice for your Home Lab.

## Features

- [x] **Spotify**

  The *Discord Service* connects to the [Spotify Service](https://github.com/quentinguidee/spotify-service) via [Redis Pub/Sub](https://redis.io/docs/manual/pubsub/), to display the artist you're currently listening, with an icon of your choice.

  <img width="231" alt="image" src="https://user-images.githubusercontent.com/12123721/219262662-e6dfaa9d-dfd6-4c7c-8e00-38e4d3c7a9ff.png">

  The status disappear when the music is paused.

- [ ] *Coming soon...*

## Setup

*This service is currently a work-in-progress.*

- Install Redis. Run an instance with the `redis-server` command.
- Install [Spotify Service](https://github.com/quentinguidee/spotify-service)
- Run `./discordservice`
