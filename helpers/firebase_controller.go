package helpers

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)
type Users struct {
	chatId string
	confirm bool
}

func AddToFirebase(collection string, data map[string]interface{}, client *firestore.Client) error {

	col := client.Collection(collection)

	// Add a new document to the collection with a generated ID.
	_, _, err := col.Add(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// fmt.Printf("Added document with ID: %v", newDocRef.ID)
	return nil
}

func GetFromFirebase(collection string, withChatId string, client *firestore.Client) (map[string]interface{}, string, error) {

	col := client.Collection(collection)

	docs, err := col.Where("chatId", "==", withChatId).Documents(context.Background()).GetAll()
	if err != nil {
        fmt.Println(err.Error())
		return nil, "", err
    }
	if len(docs) < 1 {
		return nil, "", fmt.Errorf("No Documents")
	}
	
	return docs[0].Data(), docs[0].Ref.ID, nil

}

func UpdateFirebase(collection string, doc string, field string, value interface{}, client *firestore.Client) error {

	docRef := client.Collection(collection).Doc(doc)

	_, err := docRef.Update(context.Background(), []firestore.Update{
		{Path: field, Value: value},
	})
	if err != nil {
		return err
	}
	return nil
}