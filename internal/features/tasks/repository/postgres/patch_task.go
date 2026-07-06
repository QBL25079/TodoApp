package task_postgres_repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
	core_postgres_pool "github.com/QBL25079/TodoApp/internal/core/repository/postgres/pool"
)

func (r *TaskRepository) PatchTask(ctx context.Context, taskId int, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE todoapp.tasks SET title=$1, description=$2, completed=$3, completed_at=$4, version=version + 1 WHERE id=$5 AND version=$6 RETURN id, title, description, completed, created_at, completed_at, author_user_id;`

	row := r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Completed, task.CompletedAt, taskId, task.Version)

	var taskModel TaskModel

	if err := row.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Title, &taskModel.Description, &taskModel.Completed, &taskModel.CreatedAt, &taskModel.CompletedAt, &taskModel.AuthorUserID); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("conflict: %w", core_errors.ErrConflict)
		}
		return domain.Task{}, fmt.Errorf("Scan failure: %w", err)
	}

	patchDomain := taskDomainFromModel(taskModel)
	return patchDomain, nil
}