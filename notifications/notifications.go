package notifications

func Push(title string, body string) {
	PushFirebase(title, body)
	PushExpo(title, body)
}