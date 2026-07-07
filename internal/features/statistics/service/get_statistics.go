package stat_service

import (
	"context"
	"fmt"
	"time"

	"github.com/QBL25079/TodoApp/internal/core/domain"
)

func (s *StatisticsService) GetStat(ctx context.Context, userID *int, from, to *time.Time) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf("from must be before then to")
		}
	}

	tasks, err := s.statRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("cant get array of stat structs")
	}
	stat := CalcStat(tasks)
	return stat, err
}

func CalcStat(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.Statistics{CreatedTasks: 0, CompletedTasks: 0, ComplRate: nil, AvarageTime: nil}
	}

	taskCreated := len(tasks)

	var totalCompletedDuration time.Duration

	taskComleted := 0
	for _, task := range tasks {
		if task.Completed {
			taskComleted++
		}
		complDuration := task.ComplitionDuration()
		if complDuration != nil {
			totalCompletedDuration += *complDuration
		}
	}

	taskComletedRate := float64(taskComleted) / float64(taskCreated) * 100

	var taskAvarageComletionTime *time.Duration

	if taskComleted > 0 && totalCompletedDuration != 0 {
		avg := totalCompletedDuration / time.Duration(taskComleted)
		taskAvarageComletionTime = &avg
	}

	return domain.Statistics{CreatedTasks: taskCreated, CompletedTasks: taskComleted, ComplRate: &taskComletedRate, AvarageTime: taskAvarageComletionTime}
}
