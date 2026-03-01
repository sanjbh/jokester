package agent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"go.opentelemetry.io/otel"
)

var (
	tracer   = otel.Tracer("agent")
	validate = validator.New()
)

type Agent struct {
	Name         string
	Instructions string
	Model        string
	llm          llms.Model
}

type Options struct {
	LLMBaseURL string `validate:"required,url"`
	LLMAPIKey  string `validate:"required,min=3"`
	Model      string `validate:"required"`
}

func NewAgent(name, instructions, model string) *Agent { 
		return &Agent{
		Name:         name,
		Instructions: instructions,
		Model:       
}