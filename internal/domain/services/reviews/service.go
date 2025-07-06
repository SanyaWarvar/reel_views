package reviews

import (
	"context"
	"rv/internal/domain/dto/reviews"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/trx"

	"github.com/google/uuid"
)

type reviewRepo interface {
	CreateReview(ctx context.Context, review reviews.Review) error
	GetReviewByID(ctx context.Context, id uuid.UUID) (*reviews.Review, error)
	GetReviewsByMovie(ctx context.Context, movieID uuid.UUID, limit, offset uint64) ([]reviews.Review, error)
	GetReviewsByUser(ctx context.Context, userId uuid.UUID, limit, offset uint64) ([]reviews.Review, error)
	UpdateReview(ctx context.Context, review reviews.Review) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	reviewRepo reviewRepo
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	reviewRepo reviewRepo,
) *Service {
	return &Service{
		tx:         tx,
		logger:     logger,
		reviewRepo: reviewRepo,
	}
}

func (srv *Service) CreateReview(ctx context.Context, review reviews.Review) error {
	return srv.reviewRepo.CreateReview(ctx, review)
}

func (srv *Service) GetReviewByID(ctx context.Context, id uuid.UUID) (*reviews.Review, error) {
	return srv.reviewRepo.GetReviewByID(ctx, id)
}

func (srv *Service) GetReviewsByUser(ctx context.Context, userId uuid.UUID, page uint64) ([]reviews.Review, error) {

	var offset, limit uint64
	offset = uint64(page * constants.PaginationSize)
	limit = uint64((page + 1) * constants.PaginationSize)
	return srv.reviewRepo.GetReviewsByUser(ctx, userId, limit, offset)
}

func (srv *Service) GetReviewsByMovie(ctx context.Context, movieID uuid.UUID, page uint64) ([]reviews.Review, error) {

	var offset, limit uint64
	offset = uint64(page * constants.PaginationSize)
	limit = uint64((page + 1) * constants.PaginationSize)
	return srv.reviewRepo.GetReviewsByMovie(ctx, movieID, limit, offset)
}

func (srv *Service) UpdateReview(ctx context.Context, review reviews.Review) error {
	return srv.reviewRepo.UpdateReview(ctx, review)
}

func (srv *Service) DeleteReview(ctx context.Context, id uuid.UUID) error {
	return srv.reviewRepo.DeleteReview(ctx, id)
}
