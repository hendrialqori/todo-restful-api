package web

type TodoCreateRequest struct {
	Title  string `validate:"required,max=200" json:"title"`
	UserId int    `validate:"required" json:"user_id"`
}

type TodoUpdateRequest struct {
	Id     int    `validate:"required"`
	Title  string `validate:"required,max=200" json:"title"`
	Status string `validate:"required,oneof=DONE ONPROGRESS"`
}
