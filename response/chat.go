package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
	"github.com/henomis/restclientgo"
)

type Chat struct {
	Response
	model.NonStreamedChat
	model.StreamedChat

	acceptContentType string                      `json:"-"`
	streamCallback    restclientgo.StreamCallback `json:"-"`
}

func (c *Chat) AcceptContentType() string {
	if c.acceptContentType == "" {
		return ContentTypeJSON
	}
	return c.acceptContentType
}

func (c *Chat) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}

func (c *Chat) SetAcceptContentType(acceptContentType string) {
	c.acceptContentType = acceptContentType
}

func (g *Chat) SetStreamCallback(callback restclientgo.StreamCallback) {
	g.streamCallback = callback
}

func (g *Chat) StreamCallback() restclientgo.StreamCallback {
	return g.streamCallback
}
