package user_service

import (
	"context"
	"fmt"
)

func (r *UsersService) DeleteUser(ctx context.Context, id int) error {
	if err := r.usersRepository.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}