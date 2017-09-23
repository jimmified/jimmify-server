package notifications

import (
	"jimmify-server/db"
	"github.com/adierkens/expo-server-sdk-go"
	"log"
)

var client = expo.NewExpo()

func PushExpo(title string, body string) {
	clients, err := db.GetExpoClients()
	messages := []*expo.ExpoPushMessage{}

	if err != nil {
		log.Println("Unable to fetch expo client list");
	}

	payload := make(map[string]interface{})
	payload["title"] = title
	payload["body"] = body

	for _, c := range clients {
		msg := expo.NewExpoPushMessage()
		msg.To = c.ExpoID
		msg.Title = title
		msg.Body = body
		msg.Data = payload

		messages = append(messages, msg)
	}

	_, err = client.SendPushNotifications(messages)
	if err != nil {
		log.Println(err.Error())
	}
}
