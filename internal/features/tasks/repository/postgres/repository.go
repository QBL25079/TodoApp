package task_postgres_repo

import (
	core_postgres_pool "github.com/QBL25079/TodoApp/internal/core/repository/postgres/pool"
)

type TaskRepository struct {
	pool core_postgres_pool.Pool
}

func NewTaskRepository(pool core_postgres_pool.Pool)  *TaskRepository {
	return &TaskRepository{pool: pool}
}