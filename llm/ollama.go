package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-interpreter/agent"
	"io"
	"net/http"
)

type OllamaClient struct {
	BaseURL string
	Model   string
	APIKey  string
}

func NewOllamaClient(baseURL, model, apiKey string) *OllamaClient {
	if baseURL == "" {
		baseURL = "https://ollama.com"
	}
	if model == "" {
		model = "gpt-oss:120b"
	}
	return &OllamaClient{
		BaseURL: baseURL,
		Model:   model,
		APIKey:  apiKey,
	}
}

type ollamaRequest struct {
	Model    string          `json:"model"`
	Messages []agent.Message `json:"messages"`
	Stream   bool            `json:"stream"`
}

type ollamaResponse struct {
	Message agent.Message `json:"message"`
	Done    bool          `json:"done"`
}

func (c *OllamaClient) Complete(messages []agent.Message) (string, error) {
	reqBody := ollamaRequest{
		Model:    c.Model,
		Messages: messages,
		Stream:   false,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Handle different base URL structures. 
	// If it's the cloud API, it might be just the base. The curl shows /api/chat appended.
	url := fmt.Sprintf("%s/api/chat", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ollama API error: %s", string(bodyBytes))
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return ollamaResp.Message.Content, nil
}
