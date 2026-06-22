package user_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_utils "github.com/QBL25079/TodoApp/internal/core/transport/http/request"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
)

type GetUsersResponse []DTOUserResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'limit' / 'offset' query params")
		return
	}
	userDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(UsersDTOFromDoamins(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := core_http_utils.GetQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query params: %w", err)
	}

	offset, err := core_http_utils.GetQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query params: %w", err)
	}

	return limit, offset, nil
}
