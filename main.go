package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"

	"github.com/Andrem19/gpt_chat_bot/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/option"
)
  
  func main() {
	var config helpers.Config
	var err error
	var opt option.ClientOption
	env := os.Getenv("ENV")
	if env == "production" {
		config = helpers.Config{
			GPT_BOT_TOKEN: os.Getenv("GPT_BOT_TOKEN"),
			TELEGRAM_BOT_TOKEN: os.Getenv("TELEGRAM_BOT_TOKEN"),
		}
		googleCred := os.Getenv("GOOGLE_CREDENTIALS")
		opt = option.WithCredentialsJSON([]byte(googleCred))
	} else {
		config, err = helpers.LoadConfig(".")
		opt = option.WithCredentialsFile("google-credentials.json")
	}
	

	//Start with firebase
	
	// opt = option.WithCredentialsFile("google-credentials.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "gptdb-5a185", opt)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on firebase")
	defer client.Close()
	//Start with telegram-bot
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on telegram %s", bot.Self.UserName)
	
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			go func() {
				answer, err := helpers.Switcher(update.Message.Text, strconv.FormatInt(update.Message.Chat.ID, 10), client, config.GPT_BOT_TOKEN)
				if err != nil {
					helpers.SaveError(strconv.FormatInt(update.Message.Chat.ID, 10), update.Message.Text, err.Error(), client)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}()
		}
	}
  }