package domain

import (
	"fmt"
	"time"

	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
)

type Task struct {
	Id           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func (t TaskPatch) Validate() error {
	if t.Title.Set && t.Title.Value == nil {
		return fmt.Errorf("Title value can`t be lower 0. %w",  core_errors.ErrInvalidArgument)
	}

	if t.Completed.Set && t.Completed.Value == nil {
		return fmt.Errorf("Completed value cant`be Null, %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func NewTaskPatch(title Nullable[string], description Nullable[string], completed Nullable[bool]) TaskPatch {
	return TaskPatch{Title: title, Description: description, Completed: completed}
}

func NewTask(id int, version int, title string, description *string, completed bool, createdAt time.Time, completedAt *time.Time, authorUserID int) Task {
	return Task{Id: id, Version: version, Title: title, Description: description, Completed: completed, CreatedAt: createdAt, CompletedAt: completedAt, AuthorUserID: authorUserID}
}

func NewTaskUninitialized(title string, description *string, authorUserID int) Task {
	return NewTask(UninitializedID, UninitializedVersion, title, description, false, time.Now(), nil, authorUserID)
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf("invalid title len: %d: %w", titleLen, core_errors.ErrInvalidArgument)
	}
	if t.Description != nil {
		lenDes := len([]rune(*t.Description))
		if lenDes < 1 || lenDes > 1000 {
			return fmt.Errorf("invalid description length: %d: %w", lenDes, core_errors.ErrInvalidArgument)
		}
	}
	if t.Completed == true {
		if t.CompletedAt == nil {
			return fmt.Errorf("CompletedAt can`t be nil if it already completed: %w", core_errors.ErrInvalidArgument)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("completed at cant be before created at: %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("Completed at must be nil: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task pathc: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		completed := *patch.Completed.Value

		if completed {
			CompletedAt := time.Now()
			tmp.CompletedAt = &CompletedAt
		} else {
			tmp.CompletedAt = nil
		}

		tmp.Completed = completed
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("failed to validate patched task: %w", err)
	}

	*t = tmp
	return nil
}
