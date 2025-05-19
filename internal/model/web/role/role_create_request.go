package web

type RoleCreateRequest struct {
	Name string `validate:"required,max=200,min=0" json:"name"`
}
