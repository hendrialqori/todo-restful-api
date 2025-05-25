package web

type UserCreateRequest struct {
	Email    string `validate:"required,email" json:"email"`
	UserName string `validate:"required,max=200,excludesall= " json:"username"`
	RoleId   int    `validate:"required" json:"role_id"`
}

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Email    string `validate:"required,email" json:"email"`
	UserName string `validate:"required,max=200,excludesall= " json:"username"`
}
