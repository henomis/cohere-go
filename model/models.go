package model

type Model string

const (
	ModelCommand             Model = "command"
	ModelCommandNightly      Model = "command-nightly"
	ModelCommandLight        Model = "command-light"
	ModelCommandLightNightly Model = "command-light-nightly"
	ModelCommandR            Model = "command-r"
	ModelCommandRPlus        Model = "command-r-plus"
)

func (m *Model) String() string {
	return string(*m)
}

type EmbedModel string

// English
const (
	EmbedModelEnglishV30      EmbedModel = "embed-english-v3.0"
	EmbedModelEnglishLightV30 EmbedModel = "embed-english-light-v3.0"
	EmbedModelEnglishV20      EmbedModel = "embed-english-v2.0"
	EmbedModelEnglishLightV20 EmbedModel = "embed-english-light-v2.0"
)

// Multilingual
const (
	EmbedModelMultilingualV30      EmbedModel = "embed-multilingual-v3.0"
	EmbedModelMultilingualLightV30 EmbedModel = "embed-multilingual-light-v3.0"
	EmbedModelMultilingualV20      EmbedModel = "embed-multilingual-v2.0"
)

var EmbedModelsSize = map[EmbedModel]int{
	EmbedModelEnglishV30:      1024,
	EmbedModelEnglishLightV30: 384,
	EmbedModelEnglishV20:      4096,
	EmbedModelEnglishLightV20: 1024,

	EmbedModelMultilingualV30:      1024,
	EmbedModelMultilingualLightV30: 384,
	EmbedModelMultilingualV20:      768,
}

type EmbedInputType string

const (
	EmbedSearchQuery    EmbedInputType = "search_query"
	EmbedSearchDocument EmbedInputType = "search_document"
	EmbedClassification EmbedInputType = "classification"
	EmbedClustering     EmbedInputType = "clustering"
)

type ReturnLikelihoods string

const (
	ReturnLikelihoodsGeneration ReturnLikelihoods = "GENERATION"
	ReturnLikelihoodsAll        ReturnLikelihoods = "ALL"
	ReturnLikelihoodsNone       ReturnLikelihoods = "NONE"
)

type LogitBias map[string]float64

type Truncate string

const (
	TruncateNone  Truncate = "NONE"
	TruncateStart Truncate = "START"
	TruncateEnd   Truncate = "END"
)

type Meta struct {
	APIVersion APIVersion `json:"api_version"`
	Warnings   []string   `json:"warnings"`
}

type APIVersion struct {
	Version        string `json:"version"`
	IsDeprecated   bool   `json:"is_deprecated"`
	IsExperimental bool   `json:"is_experimental"`
}

type Generation struct {
	ID               string            `json:"id"`
	Text             string            `json:"text"`
	Index            *int              `json:"index,omitempty"`
	Likelihood       *float64          `json:"likelihood,omitempty"`
	TokenLikelihoods []TokenLikelihood `json:"token_likelihoods,omitempty"`
}

type TokenLikelihood struct {
	Token      string  `json:"token"`
	Likelihood float64 `json:"likelihood"`
}

type Example struct {
	Text  string `json:"text"`
	Label string `json:"label"`
}

type Classification struct {
	ID         string              `json:"id"`
	Input      *string             `json:"input,omitempty"`
	Prediction string              `json:"prediction"`
	Confidence float64             `json:"confidence"`
	Labels     ClassificationLabel `json:"labels"`
}

type ClassificationLabel map[string]ClassificationLabelConfidence

type ClassificationLabelConfidence struct {
	Confidence float64 `json:"confidence"`
}

type DetectLanguageResult struct {
	LanguageName string `json:"language_name"`
	LanguageCode string `json:"language_code"`
}

type RerankModel string

const (
	RerankModelEnglishV20      RerankModel = "rerank-english-v2.0"
	RerankModelMultilingualV20 RerankModel = "rerank-multilingual-v2.0"
)

type RerankResult struct {
	Document       Document `json:"document"`
	Index          int64    `json:"index"`
	RelevanceScore float64  `json:"relevance_score"`
}

type Document struct {
	ID             string `json:"id"`
	Text           string `json:"text"`
	AdditionalProp any    `json:"additionalProp"`
}

type TextLength string

const (
	TextLengthShort  TextLength = "short"
	TextLengthMedium TextLength = "medium"
	TextLengthLong   TextLength = "long"
	TextLengthAuto   TextLength = "auto"
)

type TextFormat string

const (
	TextFormatParagraph TextFormat = "paragraph"
	TextFormatBullet    TextFormat = "bullet"
	TextFormatAuto      TextFormat = "auto"
)

type Extractiveness string

const (
	ExtractivenessLow    Extractiveness = "low"
	ExtractivenessMedium Extractiveness = "medium"
	ExtractivenessHigh   Extractiveness = "high"
	ExtractivenessAuto   Extractiveness = "auto"
)

type Connector struct {
	ID                string `json:"id"`
	UserAccessToken   string `json:"user_access_token,omitempty"`
	ContinueOnFailure bool   `json:"continue_on_failure,omitempty"`
	Options           any    `json:"options,omitempty"`
}

type CitationQuality string

const (
	CitationQualityAccurate CitationQuality = "accurate"
	CitationQualityFast     CitationQuality = "fast"
)

type ChatMessage struct {
	Role        ChatMessageRole `json:"role"`
	Message     string          `json:"message"`
	UserName    *string         `json:"user_name,omitempty"`
	ToolCalls   []ToolCall      `json:"tool_calls,omitempty"`
	ToolResults []ToolResult    `json:"tool_results,omitempty"`
}

type ToolCall struct {
	Name       string                 `json:"name"`
	Parameters map[string]interface{} `json:"parameters"`
}

type ToolResult struct {
	Call    ToolCall      `json:"call"`
	Outputs []interface{} `json:"outputs"`
}

type PromptTruncation string

const (
	PromptTruncationAuto              PromptTruncation = "AUTO"
	PromptTruncationOff               PromptTruncation = "OFF"
	PromptTruncationAutoPreserveOrder PromptTruncation = "AUTO_PRESERVE_ORDER"
)

type ChatMessageRole string

const (
	ChatMessageRoleUser    ChatMessageRole = "USER"
	ChatMessageRoleChatbot ChatMessageRole = "CHATBOT"
	ChatMessageRoleSystem  ChatMessageRole = "SYSTEM"
	ChatMessageRoleTool    ChatMessageRole = "TOOL"
)

type StreamedChat struct {
	EventType EventType       `json:"event_type"`
	Response  NonStreamedChat `json:"response"`
}

type FinishReason string

const (
	FinishReasonComplete   FinishReason = "COMPLETE"
	FinishReasonErrorLimit FinishReason = "ERROR_LIMIT"
	FinishReasonMaxTokens  FinishReason = "MAX_TOKENS"
	FinishReasonError      FinishReason = "ERROR"
	FinishReasonErrorToxic FinishReason = "ERROR_TOXIC"
)

type EventType string

const (
	EventTypeStreamStart             EventType = "stream-start"
	EventTypeSearchQueriesGeneration EventType = "search-queries-generation"
	EventTypeSearchResults           EventType = "search-results"
	EventTypeTextGeneration          EventType = "text-generation"
	EventTypeCitationGeneration      EventType = "citation-generation"
	EventTypeStreamEnd               EventType = "stream-end"
)

type NonStreamedChat struct {
	Text             string         `json:"text"`
	GenerationID     string         `json:"generation_id"`
	Citations        []Citation     `json:"citations"`
	Documents        []Document     `json:"documents"`
	IsSearchRequired bool           `json:"is_search_required"`
	SearchQueries    []SearchQuery  `json:"search_queries"`
	SearchResults    []SearchResult `json:"search_results"`
	FinishReason     FinishReason   `json:"finish_reason"`
	ToolCalls        []ToolCall     `json:"tool_calls"`
	ChatHistory      []ChatMessage  `json:"chat_history"`
	ForceSingleStep  bool           `json:"force_single_step,omitempty"`
}

type SearchResult struct {
	SearchQuery SearchQuery `json:"search_query"`
	Connector   Connector   `json:"connector"`
	DocumentIDs []string    `json:"document_ids"`
}

type Citation struct {
	Start       int      `json:"start"`
	End         int      `json:"end"`
	Text        string   `json:"text"`
	DocumentIDs []string `json:"document_ids"`
}

type SearchQuery struct {
	Text         string `json:"text"`
	GenerationID string `json:"generation_id"`
}

type Tool struct {
	Name                 string                             `json:"name"`
	Description          string                             `json:"description"`
	ParameterDefinitions map[string]ToolParameterDefinition `json:"parameter_definitions,omitempty"`
}

type ToolParameterDefinition struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Required    bool   `json:"required,omitempty"`
}
