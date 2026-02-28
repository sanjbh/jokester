package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	//LLM Settings
	Provider string `env:"PROVIDER" validate:"required,oneof=openai groq ollama gemini"`
	APIKey   string `env:"API_KEY" validate:"required_unless=Provider ollama,min=20"`
	BaseURL  string `env:"BASE_URL" validate:"required,url"`
	Model    string `env:"MODEL" validate:"required"`

	//App Settings
	AppEnv string `env:"APP_ENV"`
}

var validate = validator.New()

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}
	if err := validate.Struct(cfg); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			for _, v := range validationErrors {
				log.Printf(" Field: %s, Tag: %s, Value: %v\n", v.Field(), v.Tag(), v.Value())
			}
		}
		return nil, fmt.Errorf("invalid config: %v", err)
	}
	return cfg, nil
}
