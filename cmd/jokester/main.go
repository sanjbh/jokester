package main

import (
	"log"

	"github.com/sanjbh/jokester/internal/telemetry"
)

func main() {
	shutdown, err := telemetry.InitTracer("jokester")
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer shutdown()
}
