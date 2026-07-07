package stat_repository

import (
	"time"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(taskModel.ID, taskModel.Version, taskModel.Title, taskModel.Description, taskModel.Completed, taskModel.CreatedAt, taskModel.CompletedAt, taskModel.AuthorUserID)
}

func taskDomainsFromModels(tasksModels []TaskModel) []domain.Task{
	taskDomains := make([]domain.Task, len(tasksModels))

	for i, model := range tasksModels {
		taskDomains[i] = taskDomainFromModel(model)}
	return taskDomains
}