package web

type AuthLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	UserName string `validate:"required,excludesall= " json:"username"`
}
