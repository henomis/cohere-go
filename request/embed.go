package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Embed struct {
	Texts     []string              `json:"texts"`
	Model     model.EmbedModel      `json:"model,omitempty"`
	InputType *model.EmbedInputType `json:"input_type,omitempty"`
	Truncate  *model.Truncate       `json:"truncate,omitempty"`
}

func (e *Embed) Path() (string, error) {
	return "/embed", nil
}

func (e *Embed) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (e *Embed) ContentType() string {
	return "application/json"
}
