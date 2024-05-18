package models

type GeminiRequest struct {
	RequestId   string
	ResearchId  *string
	Question    string
	OutputQueue string
	Forward     *map[string]any
}
