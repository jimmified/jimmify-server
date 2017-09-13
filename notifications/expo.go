package notifications

import (
	"jimmify-server/db"
	"github.com/adierkens/expo-server-sdk-go"
	"log"
)

var client = expo.NewExpo()

func PushExpo(title string, body string) {
	q, err := db.GetExpoClients()

	log.Println(q)

	if err != nil {
		log.Println("Unable to fetch expo client list");
	}

}
