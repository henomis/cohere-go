package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Classify struct {
	Inputs   []string          `json:"inputs"`
	Examples []model.Example   `json:"examples,omitempty"`
	Model    *model.EmbedModel `json:"model,omitempty"`
	Preset   *string           `json:"preset,omitempty"`
	Truncate *model.Truncate   `json:"truncate,omitempty"`
}

func (c *Classify) Path() (string, error) {
	return "/classify", nil
}

func (c *Classify) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *Classify) ContentType() string {
	return ContentTypeJSON
}
