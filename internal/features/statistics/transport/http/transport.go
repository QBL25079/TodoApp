package stat_transport

import (
	"context"
	"net/http"
	"time"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_http_server "github.com/QBL25079/TodoApp/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticService StatisticService
}

func NewStatisticsHandler(statServ StatisticService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{statisticService: statServ}
}

type StatisticService interface {
	GetStat(ctx context.Context, userID *int, from, to *time.Time) (domain.Statistics, error)
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodGet,
			Path: "/statistics",
			Handler: h.GetStat,
		},
	}
}