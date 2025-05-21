package web

type UserResponse struct {
	Id        int          `json:"id"`
	Email     string       `json:"email"`
	UserName  string       `json:"username"`
	Role      RoleResponse `json:"role"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}
