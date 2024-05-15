package queue

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/properties"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-rabbitmq-lib/rabbitmq"
)

type queue struct {
	connection *rabbitmq.Connection
	queue      *rabbitmq.Queue
}

func (q queue) Publish(ctx context.Context, b []byte) (err error) {
	return q.queue.Publish(ctx, b)
}

func (q queue) Connect(name string) (err error) {
	q.queue = rabbitmq.NewQueue(q.connection, name, rabbitmq.CONTENT_TYPE_JSON, properties.CreateQueueIfNX())
	return q.queue.Connect()
}

func NewQueue(connection *rabbitmq.Connection) interfaces.Queue {
	return &queue{connection: connection}
}
