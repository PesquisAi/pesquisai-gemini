package main

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/client/connections"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/injector"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		panic(err)
	}
	deps := injector.NewDependencies()

	if err = connections.Connect(deps); err != nil {
		panic(err)
	}

	if err := deps.ConsumerGeminiQueue.Consume(context.Background(), deps.Controller.GeminiHandler); err != nil {
		slog.Error("Error during ai-orchestrator-callback routine: ", err)
	}

}
