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

	resp := &response.Classify{}
	err := client.Classify(
		context.Background(),
		&request.Classify{
			Inputs: []string{
				"Confirm your email address",
				"hey i need u to send some $",
			},
			Examples: []model.Example{
				{
					Text:  "Dermatologists don't like her!",
					Label: "Spam",
				},
				{
					Text:  "Hello, open to this?",
					Label: "Spam",
				},
				{
					Text:  "I need help please wire me $1000 right now",
					Label: "Spam",
				},
				{
					Text:  "Nice to know you ;)",
					Label: "Spam",
				},
				{
					Text:  "Please help me?",
					Label: "Spam",
				},
				{
					Text:  "Your parcel will be delivered today",
					Label: "Not Spam",
				},
				{
					Text:  "Review changes to our Terms and Conditions",
					Label: "Not Spam",
				},
				{
					Text:  "Weekly sync notes",
					Label: "Not Spam",
				},
				{
					Text:  "Re: Follow up from today’s meeting",
					Label: "Not Spam",
				},
				{
					Text:  "Pre-read for tomorrow",
					Label: "Not Spam",
				},
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
