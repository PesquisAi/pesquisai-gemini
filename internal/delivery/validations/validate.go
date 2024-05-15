package validations

import (
	"fmt"
	"github.com/PesquisAi/pesquisai-gemini/internal/config/errortypes"
	"github.com/PesquisAi/pesquisai-gemini/internal/delivery/dtos"
	"github.com/go-playground/validator/v10"
)

func getError(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("'%s' is required", field)
	case "min":
		return fmt.Sprintf("'%s' should be greather on lengh", field)
	case "uuid":
		return fmt.Sprintf("'%s' should be an uuid", field)
	}

	return ""
}
func getField(field string) string {
	switch field {
	case "ResearchId":
		return "research_id"
	case "RequestId":
		return "request_id"
	case "Question":
		return "question"
	case "OutputQueue":
		return "output_queue"
	case "Forward":
		return "forward"
	}
	return ""
}

func ValidateRequest(request *dtos.GeminiRequest) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(request)
	if err == nil {
		return nil
	}

	var messages []string
	for _, err := range err.(validator.ValidationErrors) {
		messages = append(messages, getError(err.ActualTag(), getField(err.Field())))
	}

	return errortypes.NewValidationException(messages...)
}
