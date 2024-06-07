package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type CheckAPIKey struct {
}

func (r *CheckAPIKey) Path() (string, error) {
	return "/check-api-key", nil
}

func (r *CheckAPIKey) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *CheckAPIKey) ContentType() string {
	return ContentTypeJSON
}
