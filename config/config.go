package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	OllamaAPIKey string
	OllamaBaseURL string
	LLMBackend    string 
}

func Load() *Config {
	_ = godotenv.Load()
	
	cfg := &Config{
		OllamaAPIKey: os.Getenv("OLLAMA_API_KEY"),
		OllamaBaseURL: os.Getenv("OLLAMA_HOST"),
		LLMBackend:    os.Getenv("LLM_BACKEND"),
	}
	
	return cfg
}
