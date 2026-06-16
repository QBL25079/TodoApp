package user_transport_http

import (
	"net/http"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_reqpuest "github.com/QBL25079/TodoApp/internal/core/transport/http/decode"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
	core_http_types "github.com/QBL25079/TodoApp/internal/core/transport/http/types"
	core_http_utils "github.com/QBL25079/TodoApp/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse DTOUserResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, err := core_http_utils.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Cant get user ID")
		return
	}
	

	var request PatchUserRequest

	if err := core_http_reqpuest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.userService.PatchUser(ctx, userID, userPatch)

	response := PatchUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
	
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{FullName: request.FullName.ToDomain(), PhoneNumber: request.PhoneNumber.ToDomain()}
}
