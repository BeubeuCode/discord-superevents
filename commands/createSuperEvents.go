package commands

import (
	"context"
	"discord-superevents/util"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/Lukaesebrot/dgc"
	"google.golang.org/api/option"
)

func CreateSuperEvent(ctx *dgc.Ctx) {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return
	}
	context := context.Background()
	sa := option.WithCredentialsFile(config.FirebaseKeyPath)
	app, err := firebase.NewApp(context, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	ctx.RespondText("")
}
