package usecases

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/models"
	"log/slog"
)

type UseCase struct {
	queue  interfaces.Queue
	gemini interfaces.Gemini
}

func (u UseCase) Execute(ctx context.Context, request models.GeminiRequest) error {
	slog.InfoContext(ctx, "useCase.Orchestrate",
		slog.String("details", "process started"))

	slog.DebugContext(ctx, "useCase.Orchestrate",
		slog.String("details", "process finished"))
	return nil
}

func NewUseCase(queue interfaces.Queue, gemini interfaces.Gemini) interfaces.UseCase {
	return &UseCase{
		queue: queue, gemini: gemini,
	}
}
