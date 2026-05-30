package user_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
)

// dto for creating user request
type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required, min=3, max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty, min=10, max=15, startswith=+"`
}

// dto for creating user response
type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("invoce CreateUser handler")
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}
	userDomain := domainFromDTO(request)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "falied to create useer")
		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{ID: user.ID, Version: user.Version, FullName: user.FullName, PhoneNumber: user.PhoneNumber}
}
