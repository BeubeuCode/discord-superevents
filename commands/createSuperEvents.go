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

func registerSuperEvent() {
	fmt.Println("not implemented")
}

// CreateSuperEvent reads the command args, creates a firebase instance and returns an ID to invoke the super event.
func CreateSuperEvent(ctx *dgc.Ctx) {

	// command arguments
	arguments := ctx.Arguments
	fmt.Println(arguments)
	fmt.Println(arguments.Amount())
	//loading config file
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return
	}

	// firestore stuff
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
	// echoing back to user
	ctx.RespondText("fin de méthode de création")
}
