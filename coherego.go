package coherego

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/henomis/cohere-go/request"
	"github.com/henomis/cohere-go/response"
	"github.com/henomis/restclientgo"
)

const (
	cohereEndpoint = "https://api.cohere.ai/v1"
)

type Client struct {
	restClient *restclientgo.RestClient
	apiKey     string
}

func New(apiKey string) *Client {
	restClient := restclientgo.New(cohereEndpoint)

	if len(apiKey) > 0 {
		restClient.SetRequestModifier(func(req *http.Request) *http.Request {
			req.Header.Set("Authorization", "Bearer "+apiKey)
			return req
		})
	}
	return &Client{
		restClient: restClient,
		apiKey:     apiKey,
	}
}

// Deprecated: this is part of legacy API
func (c *Client) Generate(ctx context.Context, req *request.Generate, res *response.Generate) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Chat(ctx context.Context, req *request.Chat, res *response.Chat) error {
	res.SetStreamCallback(nil)
	res.SetAcceptContentType(response.ContentTypeJSON)
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) ChatStream(
	ctx context.Context,
	req *request.Chat,
	res *response.Chat,
	callBackFn func(*response.Chat),
) error {
	res.SetAcceptContentType(response.ContentTypeStreamJSON)
	req.Stream = true
	res.SetStreamCallback(func(body []byte) error {
		err := json.NewDecoder(bytes.NewReader(body)).Decode(res)
		if err != nil {
			return err
		}

		callBackFn(res)
		return nil
	})
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Embed(ctx context.Context, req *request.Embed, res *response.Embed) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Classify(ctx context.Context, req *request.Classify, res *response.Classify) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Tokenize(ctx context.Context, req *request.Tokenize, res *response.Tokenize) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Detokenize(ctx context.Context, req *request.Detokenize, res *response.Detokenize) error {
	return c.restClient.Post(ctx, req, res)
}

// Deprecated: this is part of legacy API
func (c *Client) Summarize(ctx context.Context, req *request.Summarize, res *response.Summarize) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Rerank(ctx context.Context, req *request.Rerank, res *response.Rerank) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Model(ctx context.Context, req *request.Model, res *response.Model) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) ListModels(ctx context.Context, req *request.ListModels, res *response.ListModels) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) CheckAPIKey(ctx context.Context, req *request.CheckAPIKey, res *response.CheckAPIKey) error {
	return c.restClient.Post(ctx, req, res)
}
