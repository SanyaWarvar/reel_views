package apperrors

import (
	"rv/pkg/apperror"
)

var (
	InvalidAuthorizationHeader = apperror.NewUnauthorizedError("invalid authorization header", "invalid_authorization_header")
	InvalidTokenError          = apperror.NewUnauthorizedError("invalid token", "invalid_token")
)
