package response

import (
	"encoding/json"
	"io"
)

type Detokenize struct {
	Response
	Text string `json:"text"`
}

func (d *Detokenize) AcceptContentType() string {
	return "application/json"
}

func (d *Detokenize) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(d)
}
