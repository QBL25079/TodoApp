package user_transport_http

import (
	"net/http"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_utils "github.com/QBL25079/TodoApp/internal/core/transport/http/request"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user ID path value")
		return
	}

	if err := h.userService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
