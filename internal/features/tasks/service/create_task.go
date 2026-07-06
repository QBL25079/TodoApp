package task_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (t *TaskService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}

	task, err := t.taskRepository.CreateTask(ctx, task) 
	if err != nil {
		return domain.Task{}, fmt.Errorf("Cant create task: %w", err)
	}

	return task, nil
}