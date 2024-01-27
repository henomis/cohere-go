package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Chat struct {
	Response
	model.NonStreamedChat
	model.StreamedChat

	acceptContentType string `json:"-"`
}

func (g *Chat) AcceptContentType() string {
	if g.acceptContentType == "" {
		return ContentTypeJSON
	}
	return g.acceptContentType
}

func (g *Chat) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(g)
}

func (g *Chat) SetAcceptContentType(acceptContentType string) {
	g.acceptContentType = acceptContentType
}
