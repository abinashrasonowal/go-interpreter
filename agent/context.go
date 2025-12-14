package agent

import (
	"fmt"
	"runtime"
)

type Context struct {
	History []Message
}

func NewContext() *Context {
	// Detect OS and provide appropriate command examples
	var shellExamples string
	if runtime.GOOS == "windows" {
		shellExamples = `Examples:
{"tool_name": "shell", "args": {"command": "dir"}}
{"tool_name": "shell", "args": {"command": "Get-ChildItem"}}
{"tool_name": "shell", "args": {"command": "Get-Content example.txt"}}
{"tool_name": "read_file", "args": {"path": "example.txt"}}
{"tool_name": "write_file", "args": {"path": "output.txt", "content": "Hello World"}}
{"tool_name": "done", "args": {}}

Note: You are running on Windows. Use PowerShell commands (dir, Get-ChildItem, etc.).`
	} else {
		shellExamples = `Examples:
{"tool_name": "shell", "args": {"command": "ls -la"}}
{"tool_name": "shell", "args": {"command": "pwd"}}
{"tool_name": "shell", "args": {"command": "cat example.txt"}}
{"tool_name": "read_file", "args": {"path": "example.txt"}}
{"tool_name": "write_file", "args": {"path": "output.txt", "content": "Hello World"}}
{"tool_name": "done", "args": {}}

Note: You are running on Linux/Unix. Use standard Unix commands (ls, cat, pwd, etc.).`
	}

	systemPrompt := `You are a helpful Go interpreter agent. You can execute shell commands and read/write files.

Available tools:
1. shell - Execute a shell command
2. read_file - Read contents of a file
3. write_file - Write content to a file
4. done - Signal task completion

Response format: You MUST respond with valid JSON only. The 'args' field must be a flat object with string values.

` + shellExamples

	return &Context{
		History: []Message{
			{
				Role:    "system",
				Content: systemPrompt,
			},
		},
	}
}

func (c *Context) AddUserMessage(content string) {
	c.History = append(c.History, Message{Role: "user", Content: content})
}

func (c *Context) AddAssistantMessage(content string) {
	c.History = append(c.History, Message{Role: "assistant", Content: content})
}

func (c *Context) AddObservation(obs Observation) {
	content := fmt.Sprintf("Output: %s", obs.Output)
	if obs.Error != nil {
		content = fmt.Sprintf("Error: %v", obs.Error)
	}
	c.AddUserMessage(content) // Observations come as user messages in this pattern
}
