package safety

import (
	"fmt"
	"strings"
)

func ValidateCommand(cmdStr string) error {
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return fmt.Errorf("empty command")
	}
	
	baseCmd := parts[0]
	if !IsAllowedRequest(baseCmd) {
		return fmt.Errorf("command not allowed: %s", baseCmd)
	}
	return nil
}
