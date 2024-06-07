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
	Preamble         string                 `json:"preamble,omitempty"`
	ChatHistory      []model.ChatMessage    `json:"chat_history,omitempty"`
	ConversationID   string                 `json:"conversation_id,omitempty"`
	PromptTruncation model.PromptTruncation `json:"prompt_truncation,omitempty"`
	Connectors       []model.Connector      `json:"connectors,omitempty"`
	SearchQueryOnly  bool                   `json:"search_query_only,omitempty"`
	Documents        []model.Document       `json:"documents,omitempty"`
	CitationQuality  model.CitationQuality  `json:"citation_quality,omitempty"`
	Temperature      *float64               `json:"temperature,omitempty"`
	MaxTokens        *int                   `json:"max_tokens,omitempty"`
	MaxInputTokens   *int                   `json:"max_input_tokens,omitempty"`
	K                *int                   `json:"k,omitempty"`
	P                *float64               `json:"p,omitempty"`
	Seed             *float64               `json:"seed,omitempty"`
	StopSequences    []string               `json:"stop_sequences,omitempty"`
	FrequencyPenalty *float64               `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64               `json:"presence_penalty,omitempty"`
	Tools            []model.Tool           `json:"tools,omitempty"`
	ToolResults      []model.ToolResult     `json:"tool_results,omitempty"`
	ForceSingleStep  bool                   `json:"force_single_step,omitempty"`
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
