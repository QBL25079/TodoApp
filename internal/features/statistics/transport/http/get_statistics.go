package stat_transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_request "github.com/QBL25079/TodoApp/internal/core/transport/http/request"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
)

type GetStatResponse struct {
	CreatedTasks   int `json:"created_tasks"`
	CompletedTasks int	`json:"completed_tasks"`
	ComplRate      *float64 `json:"compition_rate"`
	AvarageTime    *string `json:"avavrage_time"`
}

func toDTOFromDomain(stat domain.Statistics) GetStatResponse {
	var avarageTime *string
	if stat.AvarageTime != nil {
		duration := stat.AvarageTime.String()
		avarageTime = &duration
	}
	
	return GetStatResponse{CreatedTasks: stat.CreatedTasks, CompletedTasks: stat.CompletedTasks, ComplRate: stat.ComplRate, AvarageTime: avarageTime}
}

func (h *StatisticsHTTPHandler) GetStat(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, from, to, err := getUserIDFromQueryParam(r)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get query param")
		return 
	}

	statistics, err := h.statisticService.GetStat(ctx, userId, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get stat")
		return 
	}

	response := toDTOFromDomain(statistics)

	responseHandler.JSONResponse(response, http.StatusOK)
}

type QueryParam struct {
	userID *int
	from *time.Time
	to *time.Time
}

func getUserIDFromQueryParam(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_ID"
		fromQueryParamKey  = "From"
		toQueryParamKey = "To"
	)
	
	userID, err := core_http_request.GetQueryParam(r, userIDQueryParamKey) 
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get userID query param")
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get From query param")
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get To query param")
	}

	return userID, from, to, nil
}