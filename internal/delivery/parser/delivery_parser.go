package parser

import (
	"encoding/json"
	"fmt"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/errortypes"
	"github.com/PesquisAi/pesquisai-rabbitmq-lib/rabbitmq"
)

func ParseDeliveryJSON(out interface{}, delivery amqp091.Delivery) error {
	if delivery.ContentType != rabbitmq.CONTENT_TYPE_JSON {
		return errortypes.NewValidationException(
			fmt.Sprintf("ContentType (%s) should be %s",
				delivery.ContentType, rabbitmq.CONTENT_TYPE_JSON))
	}

	return json.Unmarshal(delivery.Body, out)
}
