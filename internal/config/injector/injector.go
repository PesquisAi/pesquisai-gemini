package injector

import (
	"github.com/PesquisAi/pesquisai-gemini/internal/config/properties"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/controllers"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/usecases"
	"github.com/PesquisAi/pesquisai-rabbitmq-lib/rabbitmq"
)

type Dependencies struct {
	Controller          interfaces.Controller
	QueueConnection     *rabbitmq.Connection
	UseCase             interfaces.UseCase
	ConsumerGeminiQueue interfaces.QueueConsumer
	Queue               interfaces.Queue
}

func (d *Dependencies) Inject() *Dependencies {

	if d.QueueConnection == nil {
		d.QueueConnection = &rabbitmq.Connection{}
	}

	if d.UseCase == nil {
		d.UseCase = usecases.NewUseCase()
	}

	if d.ConsumerGeminiQueue == nil {
		queue := rabbitmq.NewQueue(
			d.QueueConnection,
			properties.QueueNameGemini,
			rabbitmq.CONTENT_TYPE_JSON,
			properties.CreateQueueIfNX())
		d.ConsumerGeminiQueue = queue
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
