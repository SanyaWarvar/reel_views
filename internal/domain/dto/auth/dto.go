package auth

import (
	"rv/internal/domain/enum"
	"time"
)

type ConfirmationCode struct {
	Code      string               `json:"code"`
	CreatedAt time.Time            `json:"createdAt"`
	Action    enum.EmailCodeAction `json:"action"`
}
