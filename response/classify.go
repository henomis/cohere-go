package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Classify struct {
	Response
	Classifications []model.Classification `json:"classifications"`
}

func (c *Classify) AcceptContentType() string {
	return "application/json"
}

func (c *Classify) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
