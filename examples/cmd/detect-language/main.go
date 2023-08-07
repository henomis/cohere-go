package main

import (
	"context"
	"fmt"
	"os"

	coherego "github.com/henomis/cohere-go"
	"github.com/henomis/cohere-go/request"
	"github.com/henomis/cohere-go/response"
)

func main() {

	client := coherego.New(os.Getenv("COHERE_API_KEY"))

	resp := &response.DetectLanguage{}
	err := client.DetectLanguage(
		context.Background(),
		&request.DetectLanguage{
			Texts: []string{
				"Hello world",
				"Ciao mondo",
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
