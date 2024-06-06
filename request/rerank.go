package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

// TODO: add rank_fields
type Rerank struct {
	Query           string             `json:"query"`
	Model           *model.RerankModel `json:"model,omitempty"`
	Documents       []string           `json:"documents"`
	TopN            *int               `json:"top_n,omitempty"`
	RankFields      []string           `json:"rank_fields,omitempty"`
	ReturnDocuments bool               `json:"return_documents"`
	MaxChunksPerDoc *int               `json:"max_chunks_per_doc,omitempty"`
}

func (r *Rerank) Path() (string, error) {
	return "/rerank", nil
}

func (r *Rerank) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *Rerank) ContentType() string {
	return ContentTypeJSON
}
