# Go Interpreter Agent

A cross-platform AI-powered agent that can execute shell commands and manage files using natural language goals.

## Features

- 🤖 Natural language goal interpretation
- 💬 Interactive chat mode (REPL)
- 🌐 Web browsing capability (read_url)
- 🖥️ Cross-platform support (Windows, Linux, macOS)
- 🔒 Built-in safety sandbox with command whitelisting
- ⏱️ Command timeout protection
- 🔄 Automatic shell detection (PowerShell on Windows, bash on Unix)

## Prerequisites

- Go 1.25.5 or higher
- Ollama running locally or accessible via API

## Installation

1. Clone the repository:
```bash
git clone https://github.com/abinashrasonowal/go-interpreter.git
cd go-interpreter
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
cp .env

# Edit .env and set your Ollama configuration
# OLLAMA_HOST=http://localhost:11434
# OLLAMA_API_KEY=
# OLLAMA_MODEL=gpt-oss:120b
```

## Building

### Windows
```powershell
go build -o agent.exe ./cmd/agent
```

### Linux/macOS
```bash
go build -o agent ./cmd/agent
```

```bash
# Start interactive mode
./agent

# Start with a single goal
./agent "list files in current directory"
```

### Examples

```bash
# List files
>> list files in current directory

# Create a file
>> create a file called test.txt with content 'Hello World'

# Read a website
>> read https://example.com and summarize it

# Run tests
>> run all tests
```

## Use Cases

### 1. Development & Quality Assurance
- **Run Tests**: `"run all tests"` (executes `go test ./...`)
- **Code Formatting**: `"format all files"` (executes `go fmt ./...`)
- **Dependency Management**: `"tidy up modules"` or `"initialize a new module"`
- **Build**: `"build the agent"`

### 2. Codebase Exploration
- **Summarize Code**: `"read agent/agent.go and explain how it works"`
- **Search Code**: `"find all places where ValidateCommand is used"`
- **Count Stats**: `"count lines of code in the actions directory"`

### 3. File Management
- **Refactoring**: `"move all markdown files to a docs folder"`
- **Cleanup**: `"remove the temp directory"`
- **Scaffolding**: `"create a new directory structure for a web server"`

### 4. Web Research
- **Fetch Info**: `"read https://pkg.go.dev/net/http and explain the Get function"`


## Architecture

- **`cmd/agent`** - Main entry point
- **`agent/`** - Core agent logic and planning
- **`actions/`** - Tool implementations (shell, file operations)
- **`llm/`** - LLM client integrations (Ollama)
- **`safety/`** - Security sandbox and command validation
- **`config/`** - Configuration management

## Safety Features

The agent includes several safety mechanisms:

1. **Command Whitelist**: Only approved commands can be executed
2. **Timeout Protection**: Commands are automatically terminated after a timeout
3. **Sandboxed Execution**: All commands run in a controlled environment

### Allowed Commands

**Unix/Linux:**
- `ls`, `cat`, `echo`, `mkdir`, `grep`, `pwd`, `touch`, `rm`, `cp`, `mv`, `find`, `wc`, `head`, `tail`

**Windows:**
- `dir`, `Get-ChildItem`, `Get-Content`

**Cross-platform:**
- `go`

## Platform-Specific Notes

### Windows
- Uses PowerShell for command execution
- The agent automatically provides PowerShell command examples to the LLM

### Linux/macOS
- Uses `/bin/sh` for command execution
- The agent automatically provides Unix command examples to the LLM

## Configuration

Environment variables (set in `.env` file):

- `OLLAMA_HOST` - Ollama API endpoint (default: http://localhost:11434)
- `OLLAMA_API_KEY` - API key for Ollama (if required)
- `OLLAMA_MODEL` - Model to use (default: gpt-oss:120b)
- `LLM_BACKEND` - LLM backend to use (default: ollama)

## License

See LICENSE file for details.