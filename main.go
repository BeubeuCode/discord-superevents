package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"benoit-arnoult.xyz/discord-superevents/util"

	"benoit-arnoult.xyz/discord-superevents/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return
	}

	// creating bot
	discord, err := discordgo.New("Bot " + config.DiscordSecretKey)
	if err != nil {
		err := fmt.Errorf("error connecting to API %s", err)
		fmt.Println(err.Error())
	}

	// We want to receive messages
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	// Open discord channel
	err = discord.Open()

	if err != nil {
		err := fmt.Errorf("error connecting to API %s", err)
		fmt.Println(err.Error())
	}

	router := dgc.Create(&dgc.Router{
		// prefixes
		Prefixes: []string{
			"!#",
		},
		// case is not important
		IgnorePrefixCase: true,
		// bots cannot trigger
		BotsAllowed: false,
		// Command array (defined later)
		Commands: []*dgc.Command{},
		// Middlwewares
		Middlewares: []dgc.Middleware{},
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong !")
		},
	})

	router.RegisterDefaultHelpCommand(discord, nil)
	// commands

	router.RegisterCmd(&dgc.Command{
		Name:        "test",
		Description: "Test command",
		Usage:       "!#test",
		Example:     "!#test",
		IgnoreCase:  true,
		Handler:     commands.TestCommand,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "create",
		Description: "Creates a SuperEvent, and returns the id of the superevent",
		Usage:       "!#create \"Title\" \"Quote\" \"Quote Author\" ImageURL \"Subtitle\" \"Description\"",
		IgnoreCase:  true,
		Handler:     commands.CreateSuperEvent,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "show",
		Description: "shows a superevent",
		Usage:       "!#show \"ID\"",
		IgnoreCase:  true,
		Handler:     commands.ShowSuperEvent,
	})

	router.Initialize(discord)

	fmt.Println("Bot is running, press CTRL-C to stop it.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}
