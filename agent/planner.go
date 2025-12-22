package agent

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type LLMClient interface {
	Complete(messages []Message) (string, error)
}

type Planner struct {
	Client LLMClient
}

func NewPlanner(client LLMClient) *Planner {
	return &Planner{Client: client}
}

func (p *Planner) Plan(ctx *Context) (*Action, error) {
	resp, err := p.Client.Complete(ctx.History)
	if err != nil {
		return nil, err
	}
	
	// Try to extract JSON from code blocks or raw text
	jsonStr := resp
	if match := regexp.MustCompile("```(?:json)?\n?([\\s\\S]*?)\n?```").FindStringSubmatch(resp); len(match) > 1 {
		jsonStr = match[1]
	} else {
		// Fallback: finding the first '{' and last '}'
		start := -1
		end := -1
		for i, r := range resp {
			if r == '{' {
				start = i
				break
			}
		}
		for i := len(resp) - 1; i >= 0; i-- {
			if resp[i] == '}' {
				end = i + 1
				break
			}
		}
		if start != -1 && end != -1 && start < end {
			jsonStr = resp[start:end]
		}
	}

	var action Action
	if err := json.Unmarshal([]byte(jsonStr), &action); err != nil {
		return nil, fmt.Errorf("failed to parse action: %v | response: %s", err, resp)
	}

	return &action, nil
}
