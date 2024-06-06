package response

import (
	"encoding/json"
	"io"
)

type Model struct {
	Response
	Name             string   `json:"name"`
	Endpoints        []string `json:"endpoints"`
	Finetuned        bool     `json:"finetuned"`
	ContextLength    int      `json:"context_length"`
	TokenizerURL     string   `json:"tokenizer_url"`
	DefaultEndpoints []string `json:"default_endpoints"`
}

func (c *Model) AcceptContentType() string {
	return ContentTypeJSON
}

func (c *Model) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
