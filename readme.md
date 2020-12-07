[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![Go Report Card](https://goreportcard.com/badge/github.com/PiiXelx64/discord-superevents)](https://goreportcard.com/report/github.com/PiiXelx64/discord-superevents)
# discord-superevents

Spice up your discord server with some superevents !

![Bot in action](bot_preview.png)
## Features

The aim of this bot is to create "superevents", inspired by HOI4 mods like TNO, or Kaiserredux.
It contains a title, a subtitle, a description, a quote, and an image.

## Getting Started

To get started you'll need to  create a [discord application](https://discord.com/developers/applications), make it a bot and get the token for the bot. You'll also need a [firebase account with a firestore db](https://firebase.google.com/docs/firestore/), and the [Authentication Json File](https://firebase.google.com/docs/admin/setup?authuser=0#use-oauth-2-0-refresh-token).

Fill in this data in the `app.env` file, which you can create following the `app.env.example` file.

### Installation/Dependencies

Download the source code, install deps
```sh
go get github.com/spf13/viper
go get firebase.google.com/go
go get github.com/bwmarrin/discordgo
go get github.com/Lukaesebrot/dgc
```
and run `go build`.
## Getting Help

An issue ? open an issue !

## License

GPL V3
