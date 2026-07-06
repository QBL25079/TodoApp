package task_postgres_repo

import (
	"context"
	"fmt"

	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
)

func (r *TaskRepository) DeleteTask(ctx context.Context, taskID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM todoapp.tasks WHERE id=$1`

	cmdTag, err := r.pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d': %w", taskID, core_errors.ErrNotFound)
	}

	return nil
}
