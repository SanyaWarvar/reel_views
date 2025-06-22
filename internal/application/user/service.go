package auth

import (
	"context"

	"rv/internal/domain/dto/request"
	userDto "rv/internal/domain/dto/user"
	"rv/pkg/applogger"
	"rv/pkg/trx"

	"github.com/google/uuid"
)

type userService interface {
	CreateUserFromAuthCredentials(ctx context.Context, credintials request.RegisterCredentials) (*userDto.User, error)
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	userService userService
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	userService userService,
) *Service {
	return &Service{
		tx:          tx,
		logger:      logger,
		userService: userService,
	}
}

func (srv *Service) RegisterUser(ctx context.Context, credentials request.RegisterCredentials) (uuid.UUID, error) {
	user, err := srv.userService.CreateUserFromAuthCredentials(ctx, credentials)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.Id, nil

}
