package builder

import (
	"encoding/json"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/models"
)

type outputQueueMessage struct {
	RequestId  string          `json:"request_id"`
	ResearchId *string         `json:"research_id,omitempty"`
	Response   string          `json:"response"`
	Forward    *map[string]any `json:"forward,omitempty"`
}

func BuildOutputQueueMessage(request models.GeminiRequest, res string) []byte {
	b, _ := json.Marshal(outputQueueMessage{
		RequestId:  request.RequestId,
		ResearchId: request.ResearchId,
		Response:   res,
		Forward:    request.Forward,
	})
	return b
}
