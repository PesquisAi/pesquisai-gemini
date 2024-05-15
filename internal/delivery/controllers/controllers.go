package controllers

import (
	"context"
	"errors"
	"github.com/PesquisAi/pesquisai-errors-lib/exceptions"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/errortypes"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/dtos"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/parser"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/validations"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-gemini/internal/domain/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type controller struct {
	useCase interfaces.UseCase
}

func (c controller) errorHandler(err error) {
	exception := &exceptions.Error{}
	if !errors.As(err, &exception) {
		exception = errortypes.NewUnknownException(err.Error())
	}

	b, _ := exception.ToJSON()
	slog.Error("controller.errorHandler",
		slog.String("details", "process error"),
		slog.String("errorType", string(b)))
}

func (c controller) def() {

}

func (c controller) GeminiHandler(delivery amqp.Delivery) {
	defer c.def()
	slog.Info("controller.GeminiHandler",
		slog.String("details", "process started"),
		slog.String("messageId", delivery.MessageId),
		slog.String("userId", delivery.UserId))

	var request dtos.GeminiRequest
	err := parser.ParseDeliveryJSON(&request, delivery)
	if err != nil {
		c.errorHandler(err)
		return
	}

	err = validations.ValidateRequest(&request)
	if err != nil {
		c.errorHandler(err)
		return
	}

	requestModel := models.GeminiRequest{
		ResearchId:  request.ResearchId,
		RequestId:   *request.RequestId,
		Question:    *request.Question,
		OutputQueue: *request.OutputQueue,
		Forward:     request.Forward,
	}

	err = c.useCase.Execute(context.Background(), requestModel)
	if err != nil {
		c.errorHandler(err)
		return
	}

	slog.Info("controller.GeminiHandler",
		slog.String("details", "process finished"))
	err = delivery.Ack(false)
	if err != nil {
		c.errorHandler(err)
		return
	}
}

func NewController(useCase interfaces.UseCase) interfaces.Controller {
	return &controller{useCase}
}
