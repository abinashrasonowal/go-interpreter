package main

import (
	"fmt"
	"os"
	"go-interpreter/actions"
	"go-interpreter/agent"
	"go-interpreter/config"
	"go-interpreter/llm"
)

func main() {
	cfg := config.Load()
	
	var client agent.LLMClient
	if cfg.OllamaAPIKey == "" {
		fmt.Println("Warning: OLLAMA_API_KEY not set")
	}

	model := "gpt-oss:120b"
	client = llm.NewOllamaClient(cfg.OllamaBaseURL, model, cfg.OllamaAPIKey)
	
	fmt.Printf("Using Ollama (Host: %s, Model: %s)\n", cfg.OllamaBaseURL, model)

	planner := agent.NewPlanner(client)
	
	ag := agent.NewAgent(planner, actions.Execute)

	if len(os.Args) < 2 {
		fmt.Println("Usage: agent <goal>")
		fmt.Println("Example: agent 'list files in current directory'")
		return
	}

	goal := os.Args[1]
	fmt.Printf("Starting agent with goal: %s\n", goal)
	ag.Run(goal)
}
