package request

type RegisterCredentials struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SendCodeRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ConfimationCodeRequest struct {
	Code  string `json:"code"  binding:"required"`
	Email string `json:"email" binding:"required"`
}
