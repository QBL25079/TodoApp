package task_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
)

func (s TaskService) GetTasks(ctx context.Context, userId, limit, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit have to be more than 0: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("Offset have to be more than 0: %w", core_errors.ErrInvalidArgument)
	}

	tasks, err := s.taskRepository.GetTasks(ctx, userId, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("Get tasks from repo: %w", err)
	}

	return tasks, nil
}
