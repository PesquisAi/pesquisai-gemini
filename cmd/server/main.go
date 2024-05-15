package main

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/connections"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/injector"
	"log/slog"
	"sync"
)

func main() {
	deps := injector.NewDependencies()

	if err := connections.Connect(deps); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		if err := deps.ConsumerGeminiQueue.Consume(context.TODO(), deps.Controller.GeminiHandler); err != nil {
			slog.Error("Error during ai-orchestrator routine: ", err)
		}
		wg.Done()
	}()

	go func() {
		if err := deps.ConsumerAiOrchestratorCallbackQueue.Consume(context.TODO(), deps.Controller.AiOrchestratorCallbackHandler); err != nil {
			slog.Error("Error during ai-orchestrator-callback routine: ", err)
		}
		wg.Done()
	}()

}
