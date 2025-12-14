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
	
	jsonStr := resp
	if match := regexp.MustCompile("```json\n([\\s\\S]*?)\n```").FindStringSubmatch(resp); len(match) > 1 {
		jsonStr = match[1]
	}

	var action Action
	if err := json.Unmarshal([]byte(jsonStr), &action); err != nil {
		return nil, fmt.Errorf("failed to parse action: %v | response: %s", err, resp)
	}

	return &action, nil
}
