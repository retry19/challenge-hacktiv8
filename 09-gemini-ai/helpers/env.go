package helpers

import "os"

var (
	GeminiBaseUrl string
	GeminiApiKey  string
	Port          string
)

func LoadEnv() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	GeminiBaseUrl = os.Getenv("GEMINI_BASE_URL")
	if GeminiBaseUrl == "" {
		panic("GEMINI_BASE_URL is required")
	}

	GeminiApiKey = os.Getenv("GEMINI_API_KEY")
	if GeminiApiKey == "" {
		panic("GEMINI_API_KEY is required")
	}
}
