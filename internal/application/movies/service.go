package movies

import (
	"context"
	"rv/internal/domain/dto/movies"
	"rv/internal/domain/dto/request"
	resp "rv/internal/domain/dto/response"
	moviesRepo "rv/internal/infrastructure/repository/movies"
	"rv/pkg/applogger"
	"rv/pkg/trx"

	"github.com/google/uuid"
)

type movieService interface {
	GetMoviesShort(ctx context.Context, filter moviesRepo.MovieFilter, page *uint64) ([]movies.MoviesShort, error)
	GetMovieFull(ctx context.Context, movieId uuid.UUID) (*movies.MoviesFull, error)
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	movieService movieService
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	movieService movieService,

) *Service {
	return &Service{
		tx:           tx,
		logger:       logger,
		movieService: movieService,
	}
}

func (srv *Service) GetMoviesShort(
	ctx context.Context,
	req request.GetMoviesShortRequest,
	host string,
) (*resp.GetMoviesShortResponse, error) {

	movies, err := srv.movieService.GetMoviesShort(ctx, moviesRepo.MovieFilter{Search: req.Search}, &req.Page)
	if err != nil {
		return nil, err
	}
	for i := range movies {
		movies[i].ImgUrl = host + "statics/images/" + movies[i].ImgUrl
	}
	return &resp.GetMoviesShortResponse{
		Movies: movies,
	}, nil
}

func (srv *Service) GetMovieFull(ctx context.Context, req request.GetMovieFullRequest, host string) (*resp.GetMovieFullResponse, error) {

	movie, err := srv.movieService.GetMovieFull(ctx, req.MovieId)
	if err != nil {
		return nil, err
	}

	movie.ImgUrl = host + "statics/images/" + movie.ImgUrl

	return &resp.GetMovieFullResponse{
		Movie: *movie,
	}, nil
}
