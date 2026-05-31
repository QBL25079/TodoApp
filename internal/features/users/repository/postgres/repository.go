package users_postgres_repo

import core_postgres_pool "github.com/QBL25079/TodoApp/internal/core/repository/postgres/pool"

type UserRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(pool core_postgres_pool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
