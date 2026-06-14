package user_transport_http

import "github.com/QBL25079/TodoApp/internal/core/domain"

// dto for creating user response
type DTOUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func UserDTOFromDomain(user domain.User) DTOUserResponse {
	return DTOUserResponse{ID: user.ID, Version: user.Version, FullName: user.FullName, PhoneNumber: user.PhoneNumber}
}

func UsersDTOFromDoamins(users []domain.User) []DTOUserResponse {
	usersDTO := make([]DTOUserResponse, len(users))

	for i, user := range users {
		usersDTO[i] = UserDTOFromDomain(user)
	}

	return usersDTO
}
