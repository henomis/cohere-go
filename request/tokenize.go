package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Tokenize struct {
	Text  string       `json:"text"`
	Model *model.Model `json:"model,omitempty"`
}

func (t *Tokenize) Path() (string, error) {
	return "/tokenize", nil
}

func (t *Tokenize) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (t *Tokenize) ContentType() string {
	return ContentTypeJSON
}
