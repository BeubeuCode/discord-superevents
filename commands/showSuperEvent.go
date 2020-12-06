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

func getSuperEventData(eventID string) (map[string]interface{}, error) {
	//loading config file
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return nil, err
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

	dsnap, err := client.Collection("superevents").Doc(eventID).Get(context)
	if err != nil {
		fmt.Printf("Error : can't get element : %s \n", err)
	}
	superEventData := dsnap.Data()
	return superEventData, nil
}

// ShowSuperEvent gets a super event from the firestore db and shows it
func ShowSuperEvent(ctx *dgc.Ctx) {
	// command arguments
	arguments := ctx.Arguments
	fmt.Println(arguments)
	fmt.Println(arguments.Amount())
	data, err := getSuperEventData(arguments.Get(0).Raw())
	if err != nil {
		fmt.Printf("ERROR GETTING SUPEREVENT DATA %s: ", err)
	}
	fmt.Printf("Document data: %#v\n", data)

	ctx.RespondText(data["Title"].(string))
	ctx.RespondText("*" + data["Subtitle"].(string) + "*")
	ctx.RespondText("**" + data["Description"].(string) + "**")
	ctx.RespondText("\" " + data["Quote"].(string) + " \"")
	ctx.RespondText(" - " + data["QuoteAuthor"].(string))
	ctx.RespondText(data["ImageURL"].(string))

}
