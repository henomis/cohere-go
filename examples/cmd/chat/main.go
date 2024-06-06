package main

import (
	"context"
	"encoding/json"
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
			Message: "Calculate 12345 + 67890",
			Tools: []model.Tool{
				{
					Name:        "sum",
					Description: "Use this tool to sum two numbers",
					ParameterDefinitions: map[string]model.ToolParameterDefinition{
						"num1": {
							Description: "The first number to sum",
							Type:        "int",
						},
						"num2": {
							Description: "The second number to sum",
							Type:        "int",
						},
					},
				},
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")

	fmt.Printf("%s\n", string(b))

}
