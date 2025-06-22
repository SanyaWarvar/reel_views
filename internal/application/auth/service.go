package auth

import (
	"context"
	"rv/internal/domain/auth/dto"
	"rv/pkg/applogger"
	"rv/pkg/trx"

	"github.com/google/uuid"
)

type Service struct {
	tx     trx.TransactionManager
	logger applogger.Logger
}

func NewAuthService(
	tx trx.TransactionManager,
	logger applogger.Logger,
) *Service {
	return &Service{
		tx:     tx,
		logger: logger,
	}
}

func (srv *Service) RegisterUser(ctx context.Context, credentials dto.AuthCredentials) (uuid.UUID, error) {
	return uuid.New(), nil
}
