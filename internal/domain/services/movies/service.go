package movies

import (
	"context"
	"rv/internal/domain/dto/movies"
	moviesRepo "rv/internal/infrastructure/repository/movies"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/trx"
)

type movieRepo interface {
	CreateMovie(ctx context.Context, movie moviesRepo.Movie) error
	GetMoviesShort(ctx context.Context, movieFilter moviesRepo.MovieFilter, offset, limit uint64) ([]movies.MoviesShort, error)
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	movieRepo movieRepo
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	movieRepo movieRepo,
) *Service {
	return &Service{
		tx:        tx,
		logger:    logger,
		movieRepo: movieRepo,
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
