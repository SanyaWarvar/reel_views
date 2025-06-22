package auth

import (
	"context"
	"rv/internal/domain/dto/request"
	dto "rv/internal/domain/dto/response"
	"rv/internal/domain/dto/user"
	userRepository "rv/internal/infrastructure/repository/user"
	"rv/pkg/applogger"
	"rv/pkg/trx"
	"time"

	"github.com/google/uuid"
)

var codeDelay time.Duration = time.Duration(time.Minute * 1)

type userService interface {
	GetUserByEmail(ctx context.Context, email string, password string) (*user.User, error)
	UpdateUser(ctx context.Context, userId uuid.UUID, filter *userRepository.UserUpdateParams) error
}

type smtpService interface {
	SendConfirmEmailCode(ctx context.Context, email string) error
	ConfirmCode(ctx context.Context, email string, code string) error
}

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger

	userService userService
	smtpService smtpService
}

func NewService(
	tx trx.TransactionManager,
	logger applogger.Logger,
	userService userService,
	smtpService smtpService,
) *Service {
	return &Service{
		tx:          tx,
		logger:      logger,
		userService: userService,
		smtpService: smtpService,
	}
}

// todo add check is confirmed email

func (srv *Service) SendConfirmationCode(ctx context.Context, req request.SendCodeRequest) (*dto.SendCodeResponse, error) {
	_, err := srv.userService.GetUserByEmail(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &dto.SendCodeResponse{NextCodeDelay: codeDelay},
		srv.smtpService.SendConfirmEmailCode(ctx, req.Email)
}

func (srv *Service) ConfirmCode(ctx context.Context, req request.ConfimationCodeRequest) error {
	u, err := srv.userService.GetUserByEmail(ctx, req.Email, "")
	if err != nil {
		return err
	}
	err = srv.smtpService.ConfirmCode(ctx, req.Email, req.Code)
	if err != nil {
		return err
	}
	t := true
	return srv.userService.UpdateUser(ctx, u.Id, &userRepository.UserUpdateParams{
		ConfirmedEmail: &t,
	})
}
