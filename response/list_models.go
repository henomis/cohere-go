package response

import (
	"encoding/json"
	"io"
)

type ListModels struct {
	Response
	Models   []Model `json:"models"`
	NextPage string  `json:"next_page"`
}

func (c *ListModels) AcceptContentType() string {
	return ContentTypeJSON
}

func (c *ListModels) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
