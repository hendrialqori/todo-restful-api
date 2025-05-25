package web

type UserResponse struct {
	Id        int          `json:"id"`
	Email     string       `json:"email"`
	UserName  string       `json:"username"`
	Role      RoleResponse `json:"role,omitzero"`
	CreatedAt string       `json:"created_at,omitempty"`
	UpdatedAt *string      `json:"updated_at,omitempty"`
}
