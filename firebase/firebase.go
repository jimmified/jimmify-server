package firebase

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

var url string
var serverKey string
var topic string

//Init : Sets the firebase key and topic from env vars
func Init() {
	url = "https://fcm.googleapis.com/fcm/send"
	serverKey = os.Getenv("JFBKEY")
	if serverKey == "" {
		log.Fatal("Missing JFBKEY for Firebase Authentication")
	}
	topic = os.Getenv("JFBTOPIC")
	if topic == "" {
		log.Fatal("Missing JFBTOPIC for Firebase Topic")
	}
}

//Push : Sends a push notification using Google Cloud Messaging
func Push(title string, body string) {
	var jsonStr = []byte(fmt.Sprintf(`{"to": "/topics/%s", "notification": {"body": "%s", "title": "%s"}}`, topic, body, title))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "key="+serverKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println("Push failed")
	}
}
