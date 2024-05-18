package interfaces

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/models"
)

type UseCase interface {
	Execute(ctx context.Context, request models.GeminiRequest) error
}
