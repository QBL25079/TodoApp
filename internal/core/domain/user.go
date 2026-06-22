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

func NewUserPatch(fullName Nullable[string], phoneNumber Nullable[string]) UserPatch {
	return UserPatch{FullName: fullName, PhoneNumber: phoneNumber}
}

func NewUser(ID, Version int, FullName string, PhoneNumber *string) User {
	return User{ID: ID, Version: Version, FullName: FullName, PhoneNumber: PhoneNumber}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return User{ID: UninitializedID, Version: UninitializedVersion, FullName: fullName, PhoneNumber: phoneNumber}
}

func (u *User) Validate() error {
	fullNameLen := len([]rune(u.FullName))

	if fullNameLen < 3 || fullNameLen > 100 {
		return fmt.Errorf("invalid fullname length %d: %w", fullNameLen, core_errors.ErrInvalidArgument)
	}
	if u.PhoneNumber != nil {
		PhoneNumberLen := len([]rune(*u.PhoneNumber))
		if PhoneNumberLen < 10 || PhoneNumberLen > 15 {
			return fmt.Errorf("invalid phone number length %d: %w", fullNameLen, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("incorrect format phone number %v: %w", re, core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("Fullname cant be changed to null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("Invald patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate updated user: %w", err)
	}
	*u = tmp

	return nil
}
