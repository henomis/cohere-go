package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type DetectLanguage struct {
	Texts []string `json:"texts"`
}

func (d *DetectLanguage) Path() (string, error) {
	return "/detect-language", nil
}

func (d *DetectLanguage) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (d *DetectLanguage) ContentType() string {
	return "application/json"
}
