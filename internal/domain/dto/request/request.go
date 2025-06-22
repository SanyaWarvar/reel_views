package request

type RegisterCredentials struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"`
}

type ConfimationCodeRequest struct {
	Code        string `json:"code"  binding:"required"`
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"newPassword"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}
