package dto

import (
	"rv/internal/domain/dto/movies"
	"rv/internal/domain/dto/reviews"
	"time"

	"github.com/google/uuid"
)

type RegisterResponse struct {
	UserId uuid.UUID `json:"userId"`
}

type SendCodeResponse struct {
	NextCodeDelay time.Duration `json:"nextCodeDelay"`
}

type ChangePictureResponse struct {
	NewImgurl string `json:"newImgUrl"`
}

type GetMoviesShortResponse struct {
	Movies []movies.MoviesShort `json:"movies"`
}

type GetMovieFullResponse struct {
	Movie movies.MoviesFull `json:"movie"`
}

type NewReviewResponse struct {
	ReviewId uuid.UUID `json:"reviewId"`
}

type ReviewListReponse struct {
	Reviews []reviews.Review `json:"reviews"`
}
