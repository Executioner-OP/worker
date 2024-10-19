package execute

import (
	"fmt"
	"os/exec"
)

// ConfigureSandbox initializes and prepares the isolate sandbox.
func ConfigureSandbox() error {
	cmd := exec.Command("isolate", "--init", "--box-id=0")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize sandbox: %v", err)
	}
	return nil
}

// CleanSandbox cleans up the sandbox after execution.
func CleanSandbox() error {
	cmd := exec.Command("isolate", "--cleanup", "--box-id=0")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clean up sandbox: %v", err)
	}
	return nil
}
