package user_service

import (
	"context"
	"fmt"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("User does not exists: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user pathch: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, id, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to save user: %w", err)
	}

	return patchedUser, nil
}
