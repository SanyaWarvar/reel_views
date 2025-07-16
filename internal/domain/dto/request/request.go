package request

import (
	"mime/multipart"
	"rv/internal/domain/dto/reviews"

	"github.com/google/uuid"
)

// RegisterCredentials
// @Schema
type RegisterCredentials struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest
// @Schema
type LoginRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"`
}

// ConfimationCodeRequest
// @Schema
type ConfimationCodeRequest struct {
	Code        string `json:"code"  binding:"required"`
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"newPassword"`
}

// ForgotPasswordRequest
// @Schema
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

// ChangeProfilePicture
// @Schema
type ChangeProfilePicture struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	UserId uuid.UUID
}

// GetMoviesShortRequest
// @Schema
type GetMoviesShortRequest struct {
	Page   uint64 `json:"page"`
	Search string `json:"search"`
}

// GetMovieFullRequest
// @Schema
type GetMovieFullRequest struct {
	MovieId uuid.UUID `json:"movieId"`
}

// NewReviewRequest
// @Schema
type NewReviewRequest struct {
	Review reviews.Review `json:"review" binding:"required"`
}

// EditReviewRequest
// @Schema
type EditReviewRequest struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
}

// DeleteReviewRequest
// @Schema
type DeleteReviewRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

// todo add validation rating

// GetMoviesShortRequest
// @Schema
type GetPersonalRecomendationsRequest struct {
	UserId uuid.UUID `json:"userId"`
}
