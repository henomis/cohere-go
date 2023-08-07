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

	resp := &response.Generate{}
	err := client.Generate(
		context.Background(),
		&request.Generate{
			Prompt: "Create a linkedin post about my new Go AI framework LinGoose.",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
