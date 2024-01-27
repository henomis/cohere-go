package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Generate struct {
	Prompt            string                   `json:"prompt"`
	Model             *model.Model             `json:"model,omitempty"`
	NumGenerations    *int                     `json:"num_generations,omitempty"`
	MaxTokens         *int                     `json:"max_tokens,omitempty"`
	Preset            *string                  `json:"preset,omitempty"`
	Temperature       *float64                 `json:"temperature,omitempty"`
	K                 *int                     `json:"k,omitempty"`
	P                 *float64                 `json:"p,omitempty"`
	FrequencyPenalty  *float64                 `json:"frequency_penalty,omitempty"`
	PresencePenalty   *float64                 `json:"presence_penalty,omitempty"`
	EndSequences      []string                 `json:"end_sequences,omitempty"`
	StopSequences     []string                 `json:"stop_sequences,omitempty"`
	ReturnLikelihoods *model.ReturnLikelihoods `json:"return_likelihoods,omitempty"`
	LogitBias         model.LogitBias          `json:"logit_bias,omitempty"`
	Truncate          *model.Truncate          `json:"truncate,omitempty"`
	Stream            bool                     `json:"stream,omitempty"`
}

func (g *Generate) Path() (string, error) {
	return "/generate", nil
}

func (g *Generate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (g *Generate) ContentType() string {
	return ContentTypeJSON
}
