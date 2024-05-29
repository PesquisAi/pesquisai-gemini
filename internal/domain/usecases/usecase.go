package usecases

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/builder"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/models"
	"log/slog"
)

type useCase struct {
	queue  interfaces.Queue
	gemini interfaces.Gemini
}

func (u useCase) Execute(ctx context.Context, request models.GeminiRequest) error {
	slog.InfoContext(ctx, "useCase.Execute",
		slog.String("details", "process started"))

	res, err := u.gemini.Ask(ctx, request.Question)
	if err != nil {
		return err
	}

	b := builder.BuildOutputQueueMessage(request, *res)

	err = u.queue.Publish(ctx, request.OutputQueue, b)
	if err != nil {
		return err
	}

	slog.DebugContext(ctx, "useCase.Execute",
		slog.String("details", "process finished"))
	return nil
}

func NewUseCase(queue interfaces.Queue, gemini interfaces.Gemini) interfaces.UseCase {
	return &useCase{
		queue: queue, gemini: gemini,
	}
}
