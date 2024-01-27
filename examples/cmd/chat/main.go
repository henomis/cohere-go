package main

import (
	"context"
	"fmt"
	"os"

	coherego "github.com/henomis/cohere-go"
	"github.com/henomis/cohere-go/model"
	"github.com/henomis/cohere-go/request"
	"github.com/henomis/cohere-go/response"
)

func main() {

	client := coherego.New(os.Getenv("COHERE_API_KEY"))

	resp := &response.Chat{}
	err := client.Chat(
		context.Background(),
		&request.Chat{
			ChatHistory: []model.ChatMessage{
				{
					Role:    model.ChatMessageRoleUser,
					Message: "Hello, how are you?",
				},
				{
					Role:    model.ChatMessageRoleChatbot,
					Message: "I'm good, how are you?",
				},
			},
			Message: "I'm very happy!",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
