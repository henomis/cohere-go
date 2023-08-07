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

	maxChunksPerDoc := 10

	resp := &response.Rerank{}
	err := client.Rerank(
		context.Background(),
		&request.Rerank{
			ReturnDocuments: true,
			MaxChunksPerDoc: &maxChunksPerDoc,
			Query:           "What is the capital of the United States?",
			Documents: []string{
				"Carson City is the capital city of the American state of Nevada.",
				"Washington, D.C. (also known as simply Washington or D.C., and officially as the District of Columbia) is the capital of the United States. It is a federal district.",
				"Capital punishment (the death penalty) has existed in the United States since beforethe United States was a country. As of 2017, capital punishment is legal in 30 of the 50 states.",
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
