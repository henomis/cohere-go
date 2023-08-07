package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Summarize struct {
	Text              string                `json:"text"`
	Length            *model.TextLength     `json:"length,omitempty"`
	Format            *model.TextFormat     `json:"format,omitempty"`
	Model             *model.Model          `json:"model,omitempty"`
	Extractiveness    *model.Extractiveness `json:"extractiveness,omitempty"`
	Temperature       *float64              `json:"temperature,omitempty"`
	AdditionalCommand *string               `json:"additional_command,omitempty"`
}

func (s *Summarize) Path() (string, error) {
	return "/summarize", nil
}

func (s *Summarize) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (s *Summarize) ContentType() string {
	return "application/json"
}
