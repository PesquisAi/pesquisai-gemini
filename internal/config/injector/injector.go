package injector

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/client/gemini"
	"github.com/PesquisAi/pesquisai-gemini/internal/client/queue"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/properties"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/controllers"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/usecases"
	"github.com/PesquisAi/pesquisai-rabbitmq-lib/rabbitmq"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Dependencies struct {
	Controller          interfaces.Controller
	QueueConnection     *rabbitmq.Connection
	UseCase             interfaces.UseCase
	ConsumerGeminiQueue interfaces.QueueConsumer
	Queue               interfaces.Queue
	Gemini              interfaces.Gemini
}

func (d *Dependencies) Inject() *Dependencies {

	if d.QueueConnection == nil {
		d.QueueConnection = &rabbitmq.Connection{}
	}

	if d.Queue == nil {
		d.Queue = queue.NewQueue(d.QueueConnection)
	}

	if d.Gemini == nil {
		keys := properties.GetGeminiApiKeys()
		clients := make([]*genai.Client, len(keys))

		for i, key := range keys {
			client, err := genai.NewClient(context.Background(), option.WithAPIKey(key))
			if err != nil {
				panic(err)
			}
			clients[i] = client
		}

		d.Gemini = gemini.NewGemini(clients...)
	}

	if d.UseCase == nil {
		d.UseCase = usecases.NewUseCase(d.Queue, d.Gemini)
	}

	if d.ConsumerGeminiQueue == nil {
		d.ConsumerGeminiQueue = rabbitmq.NewQueue(
			d.QueueConnection,
			properties.QueueNameGemini,
			rabbitmq.ContentTypeJson,
			properties.GetCreateQueueIfNX(), true, true)
	}

	if d.Controller == nil {
		d.Controller = controllers.NewController(d.UseCase)
	}
	return d
}

func NewDependencies() *Dependencies {
	deps := &Dependencies{}
	deps.Inject()
	return deps
}
