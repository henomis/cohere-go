package response

import (
	"encoding/json"
	"io"
)

type Embed struct {
	Response
	ID         string      `json:"id"`
	Embeddings [][]float64 `json:"embeddings"`
	Texts      []string    `json:"texts"`
}

func (e *Embed) AcceptContentType() string {
	return ContentTypeJSON
}

func (e *Embed) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(e)
}
