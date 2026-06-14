package users_postgres_repo

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (r *UserRepository) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, version, full_name, phone_number FROM todoapp.users ORDER BY id ASC LIMIT $1 OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Error to get users from repo: %w", err)
	}

	defer rows.Close()

	var userModels []UserModel

	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("Scan users error: %w", err)
		}
		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomain := userDomainsFromModels(userModels)

	return userDomain, nil
}