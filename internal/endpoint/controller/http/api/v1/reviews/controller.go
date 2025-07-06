package reviews

import (
	"context"
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

type reviewService interface {
	NewReview(ctx context.Context, req request.NewReviewRequest, userId uuid.UUID) (resp.NewReviewResponse, error)
	EditReview(ctx context.Context, req request.EditReviewRequest, userId uuid.UUID) error
	DeleteReview(ctx context.Context, req request.DeleteReviewRequest, userId uuid.UUID) error
	GetUserReviews(ctx context.Context, userId uuid.UUID, page uint64) (*resp.ReviewListResponse, error)
	GetMovieReviews(ctx context.Context, movieId uuid.UUID, page uint64) (*resp.ReviewListResponse, error)
}

type Controller struct {
	lgr     applogger.Logger
	builder *response.Builder

	reviewService reviewService
}

func NewController(logger applogger.Logger, builder *response.Builder, reviewService reviewService) *Controller {
	return &Controller{
		lgr:     logger,
		builder: builder,

		reviewService: reviewService,
	}
}

func (h *Controller) Init(api, authApi *gin.RouterGroup) {
	reviews := api.Group("/reviews")
	{
		reviews.GET("/movie/:page", h.getMovieReviews)
		reviews.GET("/user/:page", h.getUserReviews)
	}
	reviewsAuth := authApi.Group("/reviews")
	{

		reviewsAuth.POST("my/new", h.newReview)
		reviewsAuth.DELETE("/my", h.deleteReview)
		reviewsAuth.PUT("/my", h.editReview)
	}
}

// @Summary new_review
// @Description Добавить новую рецензию
// @Tags reviews
// @Produce json
// @Param data body request.NewReviewRequest true "data"
// @Param X-Request-Id header string true "Request id identity"
// @Param Authorization header string true "auth token"
// @Success 200 {object} response.Response{data=resp.NewReviewResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_body, invalid_X-Request-Id, invalid_authorization_header"
// @Failure 422 {object} response.Response{} "possible codes: not_unique"
// @Router /rl/api/v1/reviews/my/new [post]
func (h *Controller) newReview(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.NewReviewRequest
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperrors.InvalidAuthorizationHeader)
		return
	}
	err = c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}

	resp, err := h.reviewService.NewReview(ctx, req, userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, resp))
}

// @Summary edit_review
// @Description Отредактировать рецензию
// @Tags reviews
// @Produce json
// @Param data body request.EditReviewRequest true "data"
// @Param X-Request-Id header string true "Request id identity"
// @Param Authorization header string true "auth token"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{} "possible codes: bind_body, invalid_X-Request-Id, invalid_authorization_header"
// @Failure 422 {object} response.Response{} "possible codes: not_unique, not_my_review"
// @Router /rl/api/v1/reviews/my [put]
func (h *Controller) editReview(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.EditReviewRequest
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperrors.InvalidAuthorizationHeader)
		return
	}
	err = c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}

	err = h.reviewService.EditReview(ctx, req, userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, nil))
}

// @Summary delete_review
// @Description Удалить рецензию
// @Tags reviews
// @Produce json
// @Param data body request.DeleteReviewRequest true "data"
// @Param X-Request-Id header string true "Request id identity"
// @Param Authorization header string true "auth token"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{} "possible codes: bind_body, invalid_X-Request-Id, invalid_authorization_header"
// @Failure 422 {object} response.Response{} "possible codes: not_unique, not_my_review"
// @Router /rl/api/v1/reviews/my [delete]
func (h *Controller) deleteReview(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.DeleteReviewRequest
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperrors.InvalidAuthorizationHeader)
		return
	}
	err = c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}

	err = h.reviewService.DeleteReview(ctx, req, userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, nil))
}

// @Summary get_user_reviews
// @Description Получить рецензии, оставленные пользователем
// @Tags reviews
// @Produce json
// @Param page path int true "page"
// @Param user_id query string true "user_id"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.ReviewListResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Failure 422 {object} response.Response{} "possible codes: not_unique, not_my_review"
// @Router /rl/api/v1/reviews/user/{page} [get]
func (h *Controller) getUserReviews(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	if page == 0 {
		_ = c.Error(apperrors.ZeroPage)
		return
	}

	resp, err := h.reviewService.GetUserReviews(ctx, userId, page-1)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(200, h.builder.BuildSuccessPaginationResponse(ctx, int(page), constants.PaginationSize, 10, resp))
}

// @Summary get_movie_reviews
// @Description Получить рецензии фильма
// @Tags reviews
// @Produce json
// @Param page path int true "page"
// @Param movie_id query string true "movie_id"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=resp.ReviewListResponse}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Router /rl/api/v1/reviews/movie/{page} [get]
func (h *Controller) getMovieReviews(c *gin.Context) {
	ctx := c.Request.Context()

	movieId, err := uuid.Parse(c.Query("movie_id"))
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindPathError))
		return
	}
	if page == 0 {
		_ = c.Error(apperrors.ZeroPage)
		return
	}

	resp, err := h.reviewService.GetMovieReviews(ctx, movieId, page-1)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(200, h.builder.BuildSuccessPaginationResponse(ctx, int(page), constants.PaginationSize, 10, resp))
}
