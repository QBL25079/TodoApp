package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(ID, Version int, FullName string, PhoneNumber *string) User {
	return User{ID: ID, Version: Version, FullName: FullName, PhoneNumber: PhoneNumber}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return User{ID: UninitializedID, Version: UninitializedVersion, FullName: fullName, PhoneNumber: phoneNumber}
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid fullname length %d: %w", fullNameLength, core_errors.ErrInvalidArgument)
	}
	if u.PhoneNumber != nil {
		PhoneNumberLen := len([]rune(*u.PhoneNumber))
		if PhoneNumberLen < 10 || PhoneNumberLen > 15 {
			return fmt.Errorf("invalid phone number length %d: %w", fullNameLength, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("incorrect format phone number %v: %w", re, core_errors.ErrInvalidArgument)
		}
	}
	return nil
}
