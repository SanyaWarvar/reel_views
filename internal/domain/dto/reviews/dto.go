package reviews

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	Id          uuid.UUID `json:"id"`
	MovieId     uuid.UUID `json:"movieId" binding:"required"`
	UserId      uuid.UUID `json:"userId"`
	Description string    `json:"description" binding:"required"`
	Rating      int       `json:"rating" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}
