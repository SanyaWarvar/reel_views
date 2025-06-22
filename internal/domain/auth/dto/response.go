package dto

import "github.com/google/uuid"

type RegisterResponse struct {
	UserId uuid.UUID `json:"userId"`
}
