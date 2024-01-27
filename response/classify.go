package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Classify struct {
	Response
	ID              string                 `json:"id"`
	Classifications []model.Classification `json:"classifications"`
}

func (c *Classify) AcceptContentType() string {
	return ContentTypeJSON
}

func (c *Classify) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
