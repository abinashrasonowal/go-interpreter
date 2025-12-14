package agent

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Action represents a tool call
type Action struct {
	ToolName string            `json:"tool_name"`
	Args     map[string]string `json:"args"`
}

// Observation represents the result of an action
type Observation struct {
	Output string
	Error  error
}
