package stat_service

import (
	"context"
	"time"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

type StatisticsService struct {
	statRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(ctx context.Context, userID *int, from, to *time.Time) ([]domain.Task, error)
}

func NewStatService(statRepo StatisticsRepository) *StatisticsService {
	return &StatisticsService{statRepository: statRepo}
}
