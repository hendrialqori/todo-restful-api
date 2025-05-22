package web

type TodoResponse struct {
	Id        int          `json:"id"`
	Title     string       `json:"title"`
	Status    string       `json:"status"`
	User      UserResponse `json:"user,omitzero"`
	CreatedAt string       `json:"created_at,omitempty"`
	UpdatedAt string       `json:"updated_at,omitempty"`
}
