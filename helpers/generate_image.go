package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type RequestB struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type Response struct {
	Created int `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}
type Answer struct {
	Message string `json:"message"`
	Tokens      int64    `json:"tokens"`
}

func GenerateImage(task string, gpt_token string ) (Answer, error) {
	var err error

	client := &http.Client{}
	requestBody := RequestB{
		Prompt: task,
		N: 1,
		Size: "256x256",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	url := "https://api.openai.com/v1/images/generations"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err, req)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `Bearer `+gpt_token+``)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("The token is valid?", err)
	}
	// check if the token is valid.
	if resp.StatusCode == 401 {
		fmt.Println("The token is invalid")
		os.Exit(0)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var response Response
	json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Println(err)
	}
	builder := strings.Builder{}
	for i := 0; i < len(response.Data); i++ {
		builder.WriteString(response.Data[i].URL)
		builder.WriteString("\n")
	}
	answer := Answer{
		Message: builder.String(),
		Tokens: 65536,
	}

	return answer, nil
}