package queue

import (
	"context"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/errortypes"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/properties"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-rabbitmq-lib/rabbitmq"
)

type queue struct {
	connection *rabbitmq.Connection
	queue      *rabbitmq.Queue
}

func (q queue) Publish(ctx context.Context, name string, b []byte) (err error) {
	q.queue = rabbitmq.NewQueue(q.connection, name, rabbitmq.ContentTypeJson, properties.GetCreateQueueIfNX(), false, false)
	err = q.queue.Connect()
	if err != nil {
		return errortypes.NewQueueException(err.Error())
	}
	err = q.queue.Publish(ctx, b)
	if err != nil {
		return errortypes.NewQueueException(err.Error())
	}
	return nil
}

func NewQueue(connection *rabbitmq.Connection) interfaces.Queue {
	return &queue{connection: connection}
}
