package movies

import (
	"rv/internal/domain/dto/reviews"

	"github.com/google/uuid"
)

type MoviesShort struct {
	Id              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	ImgUrl          string    `json:"imgUrl"`
	Genres          []string  `json:"genres"`
	AvgRating       float64   `json:"avgRating"`
	SimilarityScore float64   `json:"similarityScore,omitempty"`
}

type MoviesFull struct {
	Id        uuid.UUID        `json:"id"`
	Title     string           `json:"title"`
	ImgUrl    string           `json:"imgUrl"`
	Genres    []string         `json:"genres"`
	AvgRating float64          `json:"avgRating"`
	Reviews   []reviews.Review `json:"reviews"`
}
