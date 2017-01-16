package firebase

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	"os"
)

var serverKey string
var topic string

//Init : Sets the firebase key and topic from env vars
func Init() {
	serverKey = os.Getenv("JFIREBASEKEY")
	//topic = os.Getenv("JFIREBASETOPIC")
}

//Push : Sends a push notification using Google Cloud Messaging
//msg, The message to be sent
//sum, The title of the notification
func Push(msg, sum string) {
	data := map[string]string{
		"msg": msg,
		"sum": sum,
	}

	ids := []string{
		"dW3SHCVmDsI:APA91bE-MXR-ynBpoR-CeWUrWoinA_WZE9_WjEzrgxKewlEF2r_noo840EkR-XkkoCB0FUcgAS3E96GeIFTiZUoxFVXRaALUOxE-6hLVdpS_h6HOBjfwpnwtTToLG3uRQ3HX9JAWxOB-",
	}

	c := fcm.NewFcmClient(serverKey)
	//c.NewFcmMsgTo(topic, data)
	c.NewFcmRegIdsMsg(ids, data)
	status, err := c.Send()

	fmt.Println(serverKey)
	fmt.Println(topic)

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
