package dtos

type GeminiRequest struct {
	ResearchId  *string         `json:"research_id,omitempty" validate:"omitempty,uuid"`
	RequestId   *string         `json:"request_id" validate:"required,uuid"`
	Question    *string         `json:"question" validate:"required,min=1"`
	OutputQueue *string         `json:"output_queue" validate:"required,min=1"`
	Forward     *map[string]any `json:"forward,omitempty"`
}
