package domain

import "github.com/QBL25079/TodoApp/internal/core/domain"

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(fullName string, phoneNumber *string) User {
	return User{ID: domain.UninitializedID, Version: domain.UninitializedVersion, FullName: fullName, PhoneNumber: phoneNumber}
}
