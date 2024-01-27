package response

import (
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Rerank struct {
	Response
	Results []model.RerankResult `json:"results"`
}

func (r *Rerank) AcceptContentType() string {
	return ContentTypeJSON
}

func (r *Rerank) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}
