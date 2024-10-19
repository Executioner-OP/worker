package execute

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// RunTask executes a given task in the isolate sandbox.
func RunTask(taskID, basePath string) error {
	taskPath := filepath.Join(basePath, taskID)

	// Command to execute code inside the sandbox using isolate.
	cmd := exec.Command(
		"isolate", "--box-id=0", "--run", "--env=LANG=C.UTF-8",
		"--dir=" + taskPath, // Bind task directory
		"/usr/bin/" + taskID, // Assuming binary will be built here
	)

	// Redirect the command's output and error streams to capture the results.
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("execution failed: %v\n%s", err, string(output))
	}
	fmt.Println("Execution successful:", string(output))
	return nil
}
