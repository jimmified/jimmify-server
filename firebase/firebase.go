package firebase

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	"os"
)

var serverKey string
var topic string

func init() {
	serverKey = os.Getenv("JFIREBASEKEY")
	topic = os.Getenv("JFIREBASETOPIC")
}

func push(msg, sum string) {
	data := map[string]string{
		"msg": msg,
		"sum": sum,
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
