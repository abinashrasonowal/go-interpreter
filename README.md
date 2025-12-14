# Go Interpreter Agent

A cross-platform AI-powered agent that can execute shell commands and manage files using natural language goals.

## Features

- 🤖 Natural language goal interpretation
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
git clone <repository-url>
cd go-interpreter
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
# Copy the example .env file
cp .env.example .env

# Edit .env and set your Ollama configuration
# OLLAMA_HOST=http://localhost:11434
# OLLAMA_API_KEY=
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

## Usage

### Windows
```powershell
.\agent.exe "your goal here"
```

### Linux/macOS
```bash
./agent "your goal here"
```

### Examples

```bash
# List files in current directory
./agent "list files in current directory"

# Create a file with content
./agent "create a file called test.txt with content 'Hello World'"

# Read a file
./agent "read the README.md file"

# Find all Go files
./agent "show me all .go files"
```

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
- `LLM_BACKEND` - LLM backend to use (default: ollama)

## License

See LICENSE file for details.