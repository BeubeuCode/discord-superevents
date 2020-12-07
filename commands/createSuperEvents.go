package commands

import (
	"context"
	"discord-superevents/util"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/Lukaesebrot/dgc"
	"google.golang.org/api/option"
)

//createID returns a random string with characters from the charset variable for the ID of the superevent
func createID() string {
	rand.Seed(time.Now().Unix())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var output strings.Builder
	length := 20

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charset))
		randomChar := charset[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func registerSuperEvent(Title string, Subtitle string, ImageURL string, Description string, Quote string, QuoteAuthor string) (ID string, err error) {

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
	superEventID := createID()
	fmt.Println("Event ID : " + superEventID)
	superEvent := util.SuperEvent{
		ID:          ID,
		Title:       Title,
		Description: Description,
		Quote:       Quote,
		QuoteAuthor: QuoteAuthor,
		ImageURL:    ImageURL,
		Subtitle:    Subtitle,
	}

	_, err = client.Collection("superevents").Doc(superEventID).Set(context, map[string]interface{}{
		"Title":       superEvent.Title,
		"Description": superEvent.Description,
		"Quote":       superEvent.Quote,
		"QuoteAuthor": superEvent.QuoteAuthor,
		"ImageURL":    superEvent.ImageURL,
		"Subtitle":    superEvent.Subtitle,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return superEventID, err
}

// CreateSuperEvent reads the command args, creates a firebase instance and returns an ID to invoke the super event.
func CreateSuperEvent(ctx *dgc.Ctx) {

	// command arguments
	arguments := ctx.Arguments
	fmt.Println(arguments)
	fmt.Println(arguments.Amount())
	// Fetching args into variables
	var title string = arguments.Get(0).Raw()
	var quote string = arguments.Get(1).Raw()
	var quoteAuthor string = arguments.Get(2).Raw()
	var imageURL string = arguments.Get(3).Raw()
	var subtitle string = arguments.Get(4).Raw()
	var description string = arguments.Get(5).Raw()

	ID, err := registerSuperEvent(title, subtitle, imageURL, description, quote, quoteAuthor)
	if err != nil {
		ctx.RespondText("Error creating super event. Please try again")
	} else {
		ctx.RespondText("Created super event ! ID : " + ID)
	}
	// echoing back to user
}
