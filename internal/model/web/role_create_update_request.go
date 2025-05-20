package web

type RoleCreateRequest struct {
	Name string `validate:"required,max=50,min=0" json:"name"`
}

type RoleUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,max=50,min=0" json:"name"`
}
