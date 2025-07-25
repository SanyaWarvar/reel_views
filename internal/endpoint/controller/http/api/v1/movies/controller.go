package movies

import (
	"context"
	"fmt"
	"rv/internal/domain/dto/request"
	resp "rv/internal/domain/dto/response"
	apperrors "rv/internal/errors"
	"rv/pkg/apperror"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/response"
	"rv/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type moviesService interface {
	GetMoviesShort(ctx context.Context, req request.GetMoviesShortRequest, host string) (*resp.GetMoviesShortResponse, error)
	GetMovieFull(ctx context.Context, req request.GetMovieFullRequest, host string) (*resp.GetMovieFullResponse, error)
	GetPersonalRecomedations(ctx context.Context, req request.GetPersonalRecomendationsRequest, host string) (*resp.GetMoviesShortResponse, error)
}

type Controller struct {
	lgr     applogger.Logger
	builder *response.Builder

	moviesService moviesService
}

func NewController(logger applogger.Logger, builder *response.Builder, moviesService moviesService) *Controller {
	return &Controller{
		lgr:     logger,
		builder: builder,

		moviesService: moviesService,
	}
}

func (h *Controller) Init(api, authApi *gin.RouterGroup) {
	movies := api.Group("/movies")
	moviesAuth := authApi.Group("/movies")
	{
		movies.GET("/short/:page", h.getMoviesShort)
		movies.GET("/full/:id", h.getMovie)
		moviesAuth.GET("/recomendations/user", h.getPersonalRecomedations)
	}
}

// @Summary get_movies_short
// @Description получить короткие записи о фильмах
// @Tags movies
// @Produce json
// @Param page path int true "page"
// @Param search query string false "search"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.GetMoviesShortResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id, zero_page"
// @Router /rl/api/v1/movies/short/{page} [get]
func (h *Controller) getMoviesShort(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.GetMoviesShortRequest
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	if page == 0 {
		_ = c.Error(apperrors.ZeroPage)
		return
	}

	page--
	req.Page = page

	search := c.Query("search")
	req.Search = search

	resp, err := h.moviesService.GetMoviesShort(ctx, req, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(200, h.builder.BuildSuccessPaginationResponse(ctx, int(page)+1, constants.PaginationSize, 10, resp))
}

// @Summary get_movie_full
// @Description получить полную информацию о фильме
// @Tags movies
// @Produce json
// @Param id path string true "id"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.GetMovieFullResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Router /rl/api/v1/movies/full/{id} [get]
func (h *Controller) getMovie(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.GetMovieFullRequest
	movieId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	req.MovieId = movieId
	resp, err := h.moviesService.GetMovieFull(ctx, req, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, resp))
}

// @Summary get_personal_recomendations
// @Description получить рекомендации для пользователя
// @Tags movies
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.GetMovieFullResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Router /rl/api/v1/movies/recomendations/personal/user [get]
func (h *Controller) getPersonalRecomedations(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Println(ctx.Value(constants.UserIdCtx))
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	req := request.GetPersonalRecomendationsRequest{
		UserId: userId,
	}
	resp, err := h.moviesService.GetPersonalRecomedations(ctx, req, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, resp))
}
