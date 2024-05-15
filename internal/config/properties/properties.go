package properties

import "os"

const (
	QueueNameGemini = "gemini"
	GeminiModel     = "gemini-pro"
)

func GetCreateQueueIfNX() bool {
	return os.Getenv("CREATE_QUEUE_IF_NX") == "true"
}

func GetQueueConnectionUser() string {
	return os.Getenv("QUEUE_CONNECTION_USER")
}

func GetQueueConnectionPort() string {
	return os.Getenv("QUEUE_CONNECTION_PORT")
}

func GetQueueConnectionHost() string {
	return os.Getenv("QUEUE_CONNECTION_HOST")
}

func GetQueueConnectionPassword() string {
	return os.Getenv("QUEUE_CONNECTION_PASSWORD")
}

func GetGeminiApiKey() string {
	return os.Getenv("GEMINI_API_KEY")
}
