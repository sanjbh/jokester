package main

import (
	"context"
	"log"

	"github.com/sanjbh/jokester/internal/agent"
	"github.com/sanjbh/jokester/internal/config"
	"github.com/sanjbh/jokester/internal/runner"
	"github.com/sanjbh/jokester/internal/telemetry"
	"go.opentelemetry.io/otel"
)

func main() {
	shutdown, err := telemetry.InitTracer("jokester.main")
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer shutdown()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	tracer := otel.Tracer("main")
	ctx, span := tracer.Start(context.Background(), "jokester.main")
	defer span.End()

	llmAgent, err := agent.NewAgent(
		ctx,
		"jokester", "You are a joke teller",
		&agent.Options{
			LLMBaseURL: cfg.BaseURL,
			LLMAPIKey:  cfg.APIKey,
			Model:      cfg.Model,
		},
	)
	if err != nil {
		log.Fatalf("failed to create agent: %v", err)
	}

	llmRunner := runner.New(tracer)
	result, err := llmRunner.Run(ctx, llmAgent, "tell me a joke. /no_think")
	if err != nil {
		log.Fatalf("failed to run agent: %v", err)
	}

	log.Printf("Response from LLM: %s\n", result.FinalOutput)

}
