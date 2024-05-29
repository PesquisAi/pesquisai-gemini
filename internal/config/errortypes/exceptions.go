package errortypes

import (
	"github.com/PesquisAi/pesquisai-errors-lib/exceptions"
)

const (
	UnknownCode    = "PAG001"
	ValidationCode = "PAG002"
	QueueCode      = "PAG003"
	GeminiCode     = "PAG004"
)

func NewUnknownException(message string) *exceptions.Error {
	return &exceptions.Error{Messages: []string{message}, ErrorType: exceptions.ErrorType{
		Code:  UnknownCode,
		Type:  "Unknown",
		Abort: true,
	}}
}

func NewValidationException(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  ValidationCode,
			Type:  "Validation",
			Abort: true,
		}}
}

func NewQueueException(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  QueueCode,
			Type:  "Error posting queue message",
			Abort: true,
		}}
}

func NewGeminiError(messages ...string) *exceptions.Error {
	return &exceptions.Error{
		Messages: messages,
		ErrorType: exceptions.ErrorType{
			Code:  GeminiCode,
			Type:  "Error during gemini request",
			Abort: false,
		}}
}
