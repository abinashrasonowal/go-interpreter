package actions

import (
	"fmt"
	"go-interpreter/agent"
)

var Registry = map[string]func(map[string]string) (string, error){
	"shell": func(args map[string]string) (string, error) {
		return RunShellCommand(args["command"])
	},
	"read_file": func(args map[string]string) (string, error) {
		return ReadFile(args["path"])
	},
	"write_file": func(args map[string]string) (string, error) {
		return "", WriteFile(args["path"], args["content"])
	},
}

func Execute(action agent.Action) agent.Observation {
	handler, exists := Registry[action.ToolName]
	if !exists {
		return agent.Observation{
			Error: fmt.Errorf("unknown tool: %s", action.ToolName),
		}
	}
	
	output, err := handler(action.Args)
	return agent.Observation{
		Output: output,
		Error:  err,
	}
}
