package actions

import (
	"bytes"
	"go-interpreter/safety"
	"os/exec"
	"strings"
)

func RunShellCommand(cmdStr string) (string, error) {
	if err := safety.ValidateCommand(cmdStr); err != nil {
		return "", err
	}

	ctx, cancel := safety.NewTimeoutContext()
	defer cancel()

	// Windows specific shell handling
	var cmd *exec.Cmd
	if strings.Contains(cmdStr, " ") {
		cmd = exec.CommandContext(ctx, "powershell", "-Command", cmdStr)
	} else {
		cmd = exec.CommandContext(ctx, "powershell", "-Command", cmdStr)
	}
	
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	
	if out.Len() == 0 && stderr.Len() > 0 {
		return stderr.String(), nil
	}
	
	return out.String(), nil
}
