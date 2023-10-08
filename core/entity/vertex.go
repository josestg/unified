package entity

type BisonChatRequest struct {
	Instances  []Instance `json:"instances"`
	Parameters Parameter  `json:"parameters"`
}

type BisonChatResponse struct {
	Predictions []Prediction `json:"predictions"`
	Metadata    interface{}  `json:"metadata"`
}

type BisonTextRequest struct {
	Instances  []InstanceBisonText `json:"instances"`
	Parameters Parameter           `json:"parameters"`
}

type BisonTextResponse struct {
	Predictions []Prediction `json:"predictions"`
	Metadata    interface{}  `json:"metadata"`
}

type Prediction struct {
	Candidates []Message `json:"candidates,omitempty"`
	Content    string    `json:"content,omitempty"`
}

type Parameter struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens float64 `json:"maxOutputTokens"`
	TopP            float64 `json:"topP"`
	TopK            float64 `json:"topK"`
}

type Instance struct {
	Context  string    `json:"context"`
	Examples []Example `json:"examples"`
	Messages []Message `json:"messages"`
}

type InstanceBisonText struct {
	Prompt string `json:"prompt"`
}

type Example struct {
	Input  Content `json:"input"`
	Output Content `json:"output"`
}

type Content struct {
	Content string `json:"content"`
}

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}
