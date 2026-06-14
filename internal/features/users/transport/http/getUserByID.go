package user_transport_http

import (
	"net/http"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
	core_http_utils "github.com/QBL25079/TodoApp/internal/core/transport/http/utils"
)

type GetUserResponse DTOUserResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Missing ID value")
		return
	}
	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "Cant get user with this ID")
		return
	}
	
	response := GetUserResponse(DTOUserResponse(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}