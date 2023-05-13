package helpers

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func toRussian(client *translate.Client, text string) (string, error) {
	ctx := context.Background()
	sourceLang, err := language.Parse("en")
	if err != nil {
		log.Println(err)
		return "", err
	}
	targetLang, err := language.Parse("ru")
	if err != nil {
		log.Println(err)
		return "", err
	}

	translations, err := client.Translate(ctx, []string{text}, targetLang, &translate.Options{
		Source: sourceLang,
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return translations[0].Text, nil
}

func fromRussian(client *translate.Client, text string) (string, error) {
	ctx := context.Background()
	sourceLang, err := language.Parse("ru")
	if err != nil {
		log.Println(err)
		return "", err
	}
	targetLang, err := language.Parse("en")
	if err != nil {
		log.Println(err)
		return "", err
	}

	translations, err := client.Translate(ctx, []string{text}, targetLang, &translate.Options{
		Source: sourceLang,
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return translations[0].Text, nil
}