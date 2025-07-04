package movies

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`
	ImgUrl      string         `json:"imgUrl"`
	Genres      []Genre        `json:"genres"`
	Meta        map[string]any `json:"meta"`
	CreatedAt   time.Time      `json:"createdAt"`
}

type Genre struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
