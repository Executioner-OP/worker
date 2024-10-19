package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Executioner-OP/worker/execute"
	"github.com/Executioner-OP/worker/tasks"
)

func main() {
	// Prepare directories for the task.
	basePath := "./sandbox"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		log.Fatalf("Failed to create sandbox base directory: %v", err)
	}

	// Initialize sandbox
	if err := execute.ConfigureSandbox(); err != nil {
		log.Fatalf("Failed to configure sandbox: %v", err)
	}
	defer execute.CleanSandbox()

	// Create a new task
	task := tasks.NewTask(
		"task1", "./source/main.cpp", "./input.txt", "./output.txt",
		"cpp", "g++ main.cpp -o main && ./main < input.txt > output.txt",
	)

	// Run the task inside the sandbox
	if err := execute.RunTask(task.ID, basePath); err != nil {
		log.Fatalf("Task execution failed: %v", err)
	}

	fmt.Println("Task executed successfully!")
}
