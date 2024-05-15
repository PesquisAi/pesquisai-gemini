package models

type GeminiRequest struct {
	RequestId   string          `json:"request_id"`
	ResearchId  *string         `json:"research_id"`
	Question    string          `json:"question"`
	OutputQueue string          `json:"output_queue"`
	Forward     *map[string]any `json:"forward"`
}
