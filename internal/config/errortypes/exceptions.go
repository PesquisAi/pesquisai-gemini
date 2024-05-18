package errortypes

import (
	"github.com/PesquisAi/pesquisai-errors-lib/exceptions"
)

func NewUnknownException(message string) *exceptions.Error {
	return &exceptions.Error{Messages: []string{message}, ErrorType: exceptions.ErrorType{
		Code:  "",
		Type:  "Unknown",
		Abort: true,
	}}
}

func NewValidationException(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  "",
			Type:  "Validation",
			Abort: true,
		}}
}

func NewQueueException(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  "",
			Type:  "Error posting queue message",
			Abort: true,
		}}
}

func NewGeminiError(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  "",
			Type:  "Error during gemini request",
			Abort: false,
		}}
}
