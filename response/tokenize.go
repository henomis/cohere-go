package response

import (
	"encoding/json"
	"io"
)

type Tokenize struct {
	Response
	Tokens       []int64  `json:"tokens"`
	TokenStrings []string `json:"token_strings"`
}

func (t *Tokenize) AcceptContentType() string {
	return "application/json"
}

func (t *Tokenize) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(t)
}
