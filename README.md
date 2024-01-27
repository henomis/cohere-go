# Unofficial Cohere Go SDK


[![GoDoc](https://godoc.org/github.com/henomis/cohere-go?status.svg)](https://godoc.org/github.com/henomis/cohere-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/cohere-go)](https://goreportcard.com/report/github.com/henomis/cohere-go) [![GitHub release](https://img.shields.io/github/release/henomis/cohere-go.svg)](https://github.com/henomis/cohere-go/releases)

This is [Cohere](https://cohere.com)'s **unofficial** Go client, designed to enable you to use Cohere's services easily from your own applications.

## Cohere

[Cohere](https://cohere.com) empowers every developer and enterprise to build amazing products and capture true business value with language AI.


## API support

| **Index Operations**  | **Status** |
| --- | --- |
| Generate | 🟢 | 
| Chat | 🟢 |
| Embed | 🟢 |
| Classify | 🟢 |
| Tokenize | 🟢 |
| Detokenize | 🟢 |
| Detect Language | 🟢 |
| Summarize | 🟢 |
| Rerank | 🟢 |




## Getting started

### Installation

You can load cohere-go into your project by using:
```
go get github.com/henomis/cohere-go
```


### Configuration

The only thing you need to start using Cohere's APIs is the developer API key and related environment. Copy and paste them in the corresponding place in the code, select the API and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
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
			Prompt: "Please explain to me how LLMs work",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
```

## Who uses cohere-go?

* [LinGoose](https://github.com/henomis/lingoose) Go framework for building awesome LLM apps
