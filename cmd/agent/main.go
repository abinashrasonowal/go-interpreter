package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// Parse command line args if provided
	if len(os.Args) > 1 {
		goal := os.Args[1]
		fmt.Printf("Starting agent with goal: %s\n", goal)
		ag.Run(goal)
		return
	}

	// Interactive mode
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Go Interpreter Agent (Type 'exit' or Ctrl+C to quit)")
	fmt.Println("---------------------------------------------------")

	for {
		fmt.Print("\n>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			break
		}

		if input == "" {
			continue
		}

		ag.Run(input)
	}
}
