package tasks

import (
	"fmt"
	"os"
	"path/filepath"
)

type Task struct {
	ID         string // Unique task identifier
	SourceCode string // Path to the source code file
	Input      string // Path to the input file
	Output     string // Path to the output file (where results will be written)
	LanguageID string // ID of the language (e.g., "cpp", "python")
	ExecuteCmd string // Command to execute the code in sandbox
}

// NewTask creates a new task with the provided details.
func NewTask(id, source, input, output, langID, cmd string) *Task {
	return &Task{
		ID:         id,
		SourceCode: source,
		Input:      input,
		Output:     "",
		LanguageID: langID,
		ExecuteCmd: cmd,
	}
}

// PrepareDirectories creates necessary directories for the task.
func (t *Task) PrepareDirectories(basePath string) error {
	taskPath := filepath.Join(basePath, t.ID)
	if err := os.MkdirAll(taskPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create task directory: %v", err)
	}
	return nil
}
