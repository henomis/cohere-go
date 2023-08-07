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

	resp := &response.Detokenize{}
	err := client.Detokenize(
		context.Background(),
		&request.Detokenize{
			Tokens: []int64{
				10104,
				12221,
				1315,
				34,
				1420,
				69,
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
