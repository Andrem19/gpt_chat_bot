package helpers

import (
	"time"

	"cloud.google.com/go/firestore"
)

func SaveError(chat_id string, message string, err string, client *firestore.Client) {
	data := make(map[string]interface{})
	data["error"] = err
	data["time"] = time.Now()
	data["sender"] = message
	data["msg"] = chat_id
	AddToFirebase("logs", data, client)
}