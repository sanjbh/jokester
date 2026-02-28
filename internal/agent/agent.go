package agent

import "github.com/tmc/langchaingo/llms"

type Agent struct {
	Name         string
	Instructions string
	Model        string
	llm          llms.Model
}

type Options struct {
	LLMBaseURL string
	LLMApiKey  string
	Provider   string
}

func NewAgent(name string, instructions string, model string, llm llms.Model) *Agent {
	return &Agent{
		Name:         name,
		Instructions: instructions,
		Model:        model,
		llm:          llm,
	}
}
