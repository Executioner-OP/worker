package execute

import (
	"fmt"
	"os/exec"
)

// ConfigureSandbox initializes and prepares the isolate sandbox with the given boxID.
func ConfigureSandbox(boxID string) error {
	cmd := exec.Command("isolate", "--init", fmt.Sprintf("--box-id=%s", boxID))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize sandbox with boxID %s: %v", boxID, err)
	}
	return nil
}

// CleanSandbox cleans up the sandbox with the given boxID after execution.
func CleanSandbox(boxID string) error {
	cmd := exec.Command("isolate", "--cleanup", fmt.Sprintf("--box-id=%s", boxID))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clean up sandbox with boxID %s: %v", boxID, err)
	}
	return nil
}
