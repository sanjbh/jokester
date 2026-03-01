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

func NewAgent(ctx context.Context, name, instructions string, opts *Options) (*Agent, error) {
	_, span := tracer.Start(ctx, "agent.new")
	defer span.End()

	if opts == nil {
		return nil, fmt.Errorf("options are required")
	}

	if err := validate.Struct(opts); err != nil {
		span.RecordError(err)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, v := range validationErrors {
				log.Printf(" Field: %s, Tag: %s, Value: %v\n", v.Field(), v.Tag(), v.Value())
			}
			return nil, fmt.Errorf("invalid options: %w", err)
		}
	}
	llm, err := openai.New(
		openai.WithModel(opts.Model),
		openai.WithBaseURL(opts.LLMBaseURL),
		openai.WithToken(opts.LLMAPIKey),
	)
	if err != nil {
		span.RecordError(err)
		return nil, fmt.Errorf("failed to create llm: %w", err)
	}

	return &Agent{
		Name:         name,
		Instructions: instructions,
		Model:        opts.Model,
		llm:          llm,
	}, nil
}

func (a *Agent) GetLLM() llms.Model {
	return a.llm

}
