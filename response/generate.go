package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Generate struct {
	Response
	ID          string             `json:"id"`
	Prompt      *string            `json:"prompt"`
	Generations []model.Generation `json:"generations"`
}

func (g *Generate) AcceptContentType() string {
	return "application/json"
}

func (g *Generate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(g)
}
