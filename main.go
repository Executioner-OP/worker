package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Executioner-OP/worker/execute"
	"github.com/Executioner-OP/worker/tasks"
)

func runTask(task *tasks.Task) error {
	// Prepare directories for the task.
	basePath := "./box"
	if err := task.PrepareDirectories(basePath); err != nil {
		return fmt.Errorf("failed to prepare directories for task: %v", err)
	}

	// Initialize sandbox
	if err := execute.ConfigureSandbox(task.ID); err != nil {
		return fmt.Errorf("failed to configure sandbox: %v", err)
	}

	// Write the source code and input files to the task directory
	if err := os.WriteFile(fmt.Sprintf("%s/%s/main.cpp", basePath, task.ID), []byte(task.SourceCode), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write source code file: %v", err)
	}
	if err := os.WriteFile(fmt.Sprintf("%s/%s/input.txt", basePath, task.ID), []byte(task.Input), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write input file: %v", err)
	}

	// Run the task inside the sandbox
	output, err := execute.RunTask(task.ID, basePath)
	if err != nil {
		return fmt.Errorf("task execution failed: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/%s/output.txt", basePath, task.ID), []byte(output), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	if err := execute.CleanSandbox(task.ID); err != nil {
		return fmt.Errorf("failed to clean up sandbox: %v", err)
	}

	return nil
}

func initTask() {
	// Create a new task
	task_source_code, err := os.ReadFile("./examples/main.cpp")
	if err != nil {
		log.Fatalf("Failed to read source code: %v", err)
	}
	task_input, err := os.ReadFile("./examples/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	task := tasks.NewTask(
		"task1", string(task_source_code), string(task_input), "",
		"cpp", "g++ main.cpp -o main && ./main < input.txt > output.txt",
	)

	// Run the task
	if err := runTask(task); err != nil {
		log.Fatalf("Failed to run task: %v", err)
	}

	fmt.Println("Task executed successfully!")
}

func main() {
	// Prepare directories for the task.
	// basePath := "./sandbox"
	// if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
	// 	log.Fatalf("Failed to create sandbox base directory: %v", err)
	// }

	initTask()
}
