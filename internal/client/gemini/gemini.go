package gemini

import (
	"context"
	"fmt"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/google/generative-ai-go/genai"
	"log/slog"
)

type gemini struct {
	clients []genai.Client
}

func (g gemini) Ask(ctx context.Context, input string) (result *string, err error) {
	var resp *genai.GenerateContentResponse
	for _, client := range g.clients {
		model := client.GenerativeModel("gemini-pro")

		resp, err = model.GenerateContent(ctx, genai.Text(input))
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		break
	}

	part := resp.Candidates[0].Content.Parts[0]
	str := fmt.Sprintf("%s", part)
	result = &str
	return
}

func NewGemini(clients ...genai.Client) interfaces.Gemini {
	return &gemini{clients: clients}
}
