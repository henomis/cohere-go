package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type DetectLanguage struct {
	Response
	Results []model.DetectLanguageResult `json:"results"`
}

func (d *DetectLanguage) AcceptContentType() string {
	return "application/json"
}

func (d *DetectLanguage) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(d)
}
