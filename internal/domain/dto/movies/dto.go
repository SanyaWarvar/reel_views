package movies

import "github.com/google/uuid"

type MoviesShort struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	ImgUrl    string    `json:"imgUrl"`
	Genres    []string  `json:"genres"`
	AvgRating float64   `json:"avgRating"`
}
