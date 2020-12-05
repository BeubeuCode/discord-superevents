package main

import (
	"discord-superevents/util"
	"fmt"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return
	}
	fmt.Print(config.DiscordPublicKey)
}
