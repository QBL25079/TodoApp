package domain

import "time"

type Statistics struct {
	CreatedTasks   int
	CompletedTasks int
	ComplRate      *float64
	AvarageTime    *time.Duration
}

func NewStat(created int, completed int, rate *float64, avarageTime *time.Duration) Statistics {
	return Statistics{CreatedTasks: created, CompletedTasks: completed, ComplRate: rate, AvarageTime: avarageTime}
}