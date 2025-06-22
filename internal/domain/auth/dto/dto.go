package dto

type AuthCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ConfimationCode struct {
	Code string `json:"code"`
}
