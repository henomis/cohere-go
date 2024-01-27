package request

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henomis/cohere-go/model"
)

type Chat struct {
	Message          string                 `json:"message"`
	Model            model.Model            `json:"model,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	PreambleOverride string                 `json:"preamble_override,omitempty"`
	ChatHistory      []model.ChatMessage    `json:"chat_history,omitempty"`
	ConversationID   string                 `json:"conversation_id,omitempty"`
	PromptTruncation model.PromptTruncation `json:"prompt_truncation,omitempty"`
	Connectors       []model.Connector      `json:"connectors,omitempty"`
	SearchQueryOnly  bool                   `json:"search_query_only,omitempty"`
	Documents        []model.Document       `json:"documents,omitempty"`
	CitationQuality  model.CitationQuality  `json:"citation_quality,omitempty"`
	Temperature      *float64               `json:"temperature,omitempty"`
	FrequencyPenalty *float64               `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64               `json:"presence_penalty,omitempty"`
}

func (c *Chat) Path() (string, error) {
	return "/chat", nil
}

func (c *Chat) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *Chat) ContentType() string {
	return ContentTypeJSON
}
