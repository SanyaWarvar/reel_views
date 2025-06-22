package auth

import (
	"context"

	"rv/internal/domain/dto/request"
	respDto "rv/internal/domain/dto/response"
	userDto "rv/internal/domain/dto/user"
	"rv/pkg/applogger"
	"rv/pkg/trx"
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

// todo add reg exp check for password and username and email
func (srv *Service) RegisterUser(ctx context.Context, credentials request.RegisterCredentials) (*respDto.RegisterResponse, error) {
	user, err := srv.userService.CreateUserFromAuthCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}
	return &respDto.RegisterResponse{
		UserId: user.Id,
	}, nil

}
