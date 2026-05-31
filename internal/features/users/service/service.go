package user_service

import (
	"context"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewUsersRepository(usersRepository UsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}

func NewUserService(userRepository UsersRepository) *UsersService {
	return &UsersService{usersRepository: userRepository}
}