package tasks

import (
	"errors"
	"sync"
)

type TaskManager struct {
	mu    sync.Mutex
	tasks map[string]*Task
}

// NewTaskManager initializes the task manager.
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

// AddTask adds a new task to the manager.
func (tm *TaskManager) AddTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.ID] = task
}

// GetTask retrieves a task by ID.
func (tm *TaskManager) GetTask(id string) (*Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if task, exists := tm.tasks[id]; exists {
		return task, nil
	}
	return nil, errors.New("task not found")
}
