package interfaces

import amqp "github.com/rabbitmq/amqp091-go"

type Controller interface {
	GeminiHandler(delivery amqp.Delivery) error
}
