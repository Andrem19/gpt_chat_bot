package helpers

import (
	"fmt"

	"cloud.google.com/go/firestore"
)

func Switcher(message string, chat_id string, client *firestore.Client, GPT_BOT_TOKEN string) (string, error) {

	user, id, err := GetFromFirebase("users", chat_id, client)
	if err != nil {
		if err.Error() == "No Documents" {
			data := make(map[string]interface{})
			data["chatId"] = chat_id
			data["confirmed"] = false
			data["tokensUsed"] = 0
			AddToFirebase("users", data, client)
			return "Your chatId was successfuly registred. Ask Andrew to confirm your access.\nIf you dont know who is Andrew, this chat is not for you", nil
		}
		return "", err
	}
	rawValue, ok := user["confirmed"]
	if !ok {
		return "User does not exist", nil
	}
	value, ok := rawValue.(bool)
	if !ok {
		return "Server error", nil
	}

	isConfirmed := value
	if !isConfirmed {
		return "Ask Andrew to confirm your access", nil
	}
	if len(message) > 0 {
		rawValue, ok := user["tokensUsed"]
		if !ok {
			return "Something went wrong 2", nil
		}
		tokens, ok := rawValue.(int64)
		if !ok {
			return "Something went wrong 3", nil
		}
		if message[0:2] == "-i" {
			answer, err := GenerateImage(message, GPT_BOT_TOKEN)
			if err != nil {
				SaveError(chat_id, message, err.Error(), client)
			}
			
			UpdateFirebase("users", id, "tokensUsed", tokens + answer.Tokens, client)
			return answer.Message, nil
		} else if message[0:2] == "-c" {
			commands := Decode(message)
			return CountPriceAndAmounts(commands)
		} else {
			answer, err := AskQuestion(message, GPT_BOT_TOKEN)
			if err != nil {
				SaveError(chat_id, message, err.Error(), client)
			}
			UpdateFirebase("users", id, "tokensUsed", tokens + answer.Tokens, client)
			return answer.Message, nil
		}

	}
	return "", fmt.Errorf("Something went wrong 1")
}
