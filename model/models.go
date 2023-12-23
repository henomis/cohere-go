package model

type Model string

const (
	ModelCommand             Model = "command"
	ModelCommandNightly      Model = "command-nightly"
	ModelCommandLight        Model = "command-light"
	ModelCommandLightNightly Model = "command-light-nightly"
)

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
	ApiVersion ApiVersion `json:"api_version"`
	Warnings   []string   `json:"warnings"`
}

type ApiVersion struct {
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
	Text string `json:"text"`
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
