package main

import (
	"discord-superevents/util"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
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
		err := fmt.Errorf("Error Connecting to API %s", err)
		fmt.Println(err.Error())
	}

	// setting the message handler
	discord.AddHandler(messageCreate)
	// We want to receive messages
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	// Open discord channel
	err = discord.Open()

	if err != nil {
		err := fmt.Errorf("Error Connecting to API %s", err)
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running, press CTRL-C to stop it.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// preventing the bot from replying to itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong !")
	}
}
