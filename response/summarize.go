package response

import (
	"encoding/json"
	"io"
)

type Summarize struct {
	Response
	ID      string `json:"id"`
	Summary string `json:"summary"`
}

func (s *Summarize) AcceptContentType() string {
	return ContentTypeJSON
}

func (s *Summarize) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(s)
}
