package connections

import (
	"github.com/PesquisAi/pesquisai-gemini/internal/config/injector"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/properties"
)

func Connect(deps *injector.Dependencies) error {

	err := deps.QueueConnection.Connect(
		properties.QueueConnectionUser(),
		properties.QueueConnectionPassword(),
		properties.QueueConnectionHost(),
		properties.QueueConnectionPort(),
	)
	if err != nil {
		return err
	}

	err = deps.ConsumerGeminiQueue.Connect()
	if err != nil {
		return err
	}

	return nil
}
