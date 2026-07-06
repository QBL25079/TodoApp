package task_service

import (
	"context"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

type TaskService struct {
	taskRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userId, limit, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
	PatchTask(ctx context.Context, taskId int, task domain.Task) (domain.Task, error)
}

func NewTaskService(taskRepository TasksRepository) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}
