package task_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (s *TaskService) GetTask(ctx context.Context, taskID int) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	return task, nil
}