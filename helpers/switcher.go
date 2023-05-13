package helpers

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/translate"
)

func Switcher(message string, chat_id string, client *firestore.Client, config *Config, clientTranslate *translate.Client) (string, error) {

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
			msg, err := fromRussian(clientTranslate, message[2:])
			if err != nil {
				return "", err
			}
			answer, err := GenerateImage(msg, config.GPT_BOT_TOKEN)
			if err != nil {
				SaveError(chat_id, message, err.Error(), client)
			}
			
			UpdateFirebase("users", id, "tokensUsed", tokens + answer.Tokens, client)
			return answer.Message, nil
		} else if message[0:2] == "-c" {
			commands := Decode(message)
			return CountPriceAndAmounts(commands)
		} else {
			msg, err := fromRussian(clientTranslate, message)
			if err != nil {
				return "", err
			}
			answer, err := AskQuestion(msg, config.GPT_BOT_TOKEN, clientTranslate)
			if err != nil {
				SaveError(chat_id, message, err.Error(), client)
			}
			UpdateFirebase("users", id, "tokensUsed", tokens + answer.Tokens, client)
			return answer.Message, nil
		}

	}
	return "", fmt.Errorf("Something went wrong 1")
}
