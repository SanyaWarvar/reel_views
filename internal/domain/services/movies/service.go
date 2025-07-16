package movies

import (
	"context"
	"rv/internal/domain/dto/movies"
	"rv/internal/domain/dto/reviews"
	apperrors "rv/internal/errors"
	moviesRepo "rv/internal/infrastructure/repository/movies"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/trx"

	"github.com/google/uuid"
)

type movieRepo interface {
	CreateMovie(ctx context.Context, movie moviesRepo.Movie) error
	GetMoviesShort(ctx context.Context, movieFilter moviesRepo.MovieFilter, offset, limit uint64) ([]movies.MoviesShort, error)
	GetRecomendationsForMovie(ctx context.Context, movieId uuid.UUID) ([]movies.MoviesShort, error)
	GetRecomendationsForUser(ctx context.Context, userId uuid.UUID) ([]movies.MoviesShort, error)
}

type reviewRepo interface {
	GetReviewsByMovie(ctx context.Context, movieID uuid.UUID, limit, offset uint64) ([]reviews.Review, error)
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	movieRepo  movieRepo
	reviewRepo reviewRepo
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	movieRepo movieRepo,
	reviewRepo reviewRepo,
) *Service {
	return &Service{
		tx:         tx,
		logger:     logger,
		movieRepo:  movieRepo,
		reviewRepo: reviewRepo,
	}
}

func (srv *Service) GetMoviesShort(ctx context.Context, filter moviesRepo.MovieFilter, page *uint64) ([]movies.MoviesShort, error) {
	var offset, limit uint64
	if page == nil {
		offset = 0
		limit = 1_000_000_000
	} else {
		offset = uint64(*page * constants.PaginationSize)
		limit = uint64((*page + 1) * constants.PaginationSize)
	}
	return srv.movieRepo.GetMoviesShort(ctx, filter, offset, limit)
}

func (srv *Service) GetMovieFull(ctx context.Context, movieId uuid.UUID) (*movies.MoviesFull, error) {
	short, err := srv.movieRepo.GetMoviesShort(ctx, moviesRepo.MovieFilter{Id: &movieId}, 0, 1)
	if err != nil {
		return nil, apperrors.MovieNotFound
	}
	current := short[0]
	reviews, err := srv.reviewRepo.GetReviewsByMovie(ctx, movieId, 50, 0)
	return &movies.MoviesFull{
		Id:        current.Id,
		Title:     current.Title,
		ImgUrl:    current.ImgUrl,
		Genres:    current.Genres,
		AvgRating: current.AvgRating,
		Reviews:   reviews,
	}, err
}

func (srv *Service) GetRecomendationsForMovie(ctx context.Context, movieId uuid.UUID) ([]movies.MoviesShort, error) {
	return srv.movieRepo.GetRecomendationsForMovie(ctx, movieId)
}

func (srv *Service) GetRecomendationsForUser(ctx context.Context, userId uuid.UUID) ([]movies.MoviesShort, error) {
	return srv.movieRepo.GetRecomendationsForUser(ctx, userId)
}
