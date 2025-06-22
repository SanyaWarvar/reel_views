package apperrors

import (
	"rv/pkg/apperror"
)

var (
	InvalidAuthorizationHeader = apperror.NewUnauthorizedError("invalid authorization header", "invalid_authorization_header")
	InvalidTokenError          = apperror.NewUnauthorizedError("invalid token", "invalid_token")

	UserNotFound           = apperror.NewInvalidDataError("user not found", "user_not_found")
	IncorrectPassword      = apperror.NewInvalidDataError("incorrect password", "incorrect_password")
	ConfirmCodeAlreadySend = apperror.NewInvalidDataError("confirm code already send", "confirm_code_already_send")
	ConfirmCodeNotExist    = apperror.NewInvalidDataError("confirm code not exist", "confirm_code_not_exist")
	ConfirmCodeIncorrect   = apperror.NewInvalidDataError("confirm code incorrect", "confirm_code_incorrect")
)
