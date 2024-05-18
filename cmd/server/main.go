package main

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/client/connections"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/injector"
	"log/slog"
)

func main() {
	deps := injector.NewDependencies()

	if err := connections.Connect(deps); err != nil {
		panic(err)
	}

	if err := deps.ConsumerGeminiQueue.Consume(context.TODO(), deps.Controller.GeminiHandler); err != nil {
		slog.Error("Error during ai-orchestrator routine: ", err)
	}

}
