package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Detokenize struct {
	Tokens []int64      `json:"tokens"`
	Model  *model.Model `json:"model,omitempty"`
}

func (d *Detokenize) Path() (string, error) {
	return "/detokenize", nil
}

func (d *Detokenize) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (d *Detokenize) ContentType() string {
	return "application/json"
}
