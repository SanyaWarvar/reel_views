package movies

import (
	"context"
	"rv/internal/domain/dto/request"
	resp "rv/internal/domain/dto/response"
	"rv/pkg/apperror"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type moviesService interface {
	GetMoviesShort(ctx context.Context, req request.GetMoviesShortRequest, host string) (*resp.GetMoviesShortResponse, error)
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
	//moviesAuth := authApi.Group("/movies")
	{
		movies.GET("/short/:page", h.getMoviesShort)
	}
}

// @Summary get_movies_short
// @Description получить короткие записи о фильмах
// @Tags user
// @Produce json
// @Param data body request.ChangeProfilePicture true "data"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.GetMoviesShortResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Router /rl/api/v1/movies/short/{page} [get]
func (h *Controller) getMoviesShort(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.GetMoviesShortRequest
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}

	req.Page = page
	resp, err := h.moviesService.GetMoviesShort(ctx, req, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(200, h.builder.BuildSuccessPaginationResponse(ctx, int(page), constants.PaginationSize, 1000, resp))
}

// todo all pages
