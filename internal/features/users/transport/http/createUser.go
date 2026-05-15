package user_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
)

// dto for creating user request
type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
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
	log.Debug("invoce CreateUser handler")
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}
}
