package reviews

import (
	"context"
	"rv/internal/domain/dto/request"
	"rv/internal/domain/dto/reviews"
	apperrors "rv/internal/errors"
	"rv/pkg/applogger"
	"rv/pkg/trx"
	"rv/pkg/util"

	resp "rv/internal/domain/dto/response"

	"github.com/google/uuid"
)

type reviewSrv interface {
	CreateReview(ctx context.Context, review reviews.Review) error
	GetReviewByID(ctx context.Context, id uuid.UUID) (*reviews.Review, error)
	GetReviewsByMovie(ctx context.Context, movieID uuid.UUID, page uint64) ([]reviews.Review, error)
	GetReviewsByUser(ctx context.Context, userId uuid.UUID, page uint64) ([]reviews.Review, error)
	UpdateReview(ctx context.Context, review reviews.Review) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	reviewSrv reviewSrv
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	reviewSrv reviewSrv,

) *Service {
	return &Service{
		tx:        tx,
		logger:    logger,
		reviewSrv: reviewSrv,
	}
}

func (srv *Service) NewReview(ctx context.Context, req request.NewReviewRequest, userId uuid.UUID) (resp.NewReviewResponse, error) {
	req.Review.UserId = userId
	req.Review.Id = util.NewUUID()
	req.Review.CreatedAt = util.GetCurrentUTCTime()
	err := srv.reviewSrv.CreateReview(ctx, req.Review)
	return resp.NewReviewResponse{
		ReviewId: req.Review.Id,
	}, err
}

func (srv *Service) EditReview(ctx context.Context, req request.EditReviewRequest, userId uuid.UUID) error {
	review, err := srv.reviewSrv.GetReviewByID(ctx, req.Id)
	if err != nil {
		return apperrors.ReviewNotFound
	}
	if review.UserId != userId {
		return apperrors.NotMyReview
	}
	if req.Description != "" {
		review.Description = req.Description
	}
	if req.Rating != 0 {
		review.Rating = req.Rating
	}
	return srv.reviewSrv.UpdateReview(ctx, *review)
}

func (srv *Service) DeleteReview(ctx context.Context, req request.DeliteReviewRequest, userId uuid.UUID) error {
	review, err := srv.reviewSrv.GetReviewByID(ctx, req.Id)
	if err != nil {
		return apperrors.ReviewNotFound
	}
	if review.UserId != userId {
		return apperrors.NotMyReview
	}

	return srv.reviewSrv.DeleteReview(ctx, req.Id)
}

func (srv *Service) GetUserReviews(ctx context.Context, userId uuid.UUID, page uint64) (*resp.ReviewListReponse, error) {
	reviews, err := srv.reviewSrv.GetReviewsByUser(ctx, userId, page)
	return &resp.ReviewListReponse{Reviews: reviews}, err
}

func (srv *Service) GetMovieReviews(ctx context.Context, movieId uuid.UUID, page uint64) (*resp.ReviewListReponse, error) {
	reviews, err := srv.reviewSrv.GetReviewsByMovie(ctx, movieId, page)
	return &resp.ReviewListReponse{Reviews: reviews}, err
}
