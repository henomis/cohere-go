package coherego

import (
	"context"
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

func (c *Client) Generate(ctx context.Context, req *request.Generate, res *response.Generate) error {
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

func (c *Client) DetectLanguage(ctx context.Context, req *request.DetectLanguage, res *response.DetectLanguage) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) Rerank(ctx context.Context, req *request.Rerank, res *response.Rerank) error {
	return c.restClient.Post(ctx, req, res)
}
