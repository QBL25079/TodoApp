package task_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (s *TaskService) PatchTask(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, id) 
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("Apply task patch: %w", err)
	}

	patchedTask, err := s.taskRepository.PatchTask(ctx, id, task)

	if err != nil {
		return domain.Task{}, fmt.Errorf("Failed to update task: %w", err)
	}

	return patchedTask, nil
}