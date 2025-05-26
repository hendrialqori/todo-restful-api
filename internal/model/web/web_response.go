package web

type ApiResponse[T any] struct {
	Ok      bool   `json:"ok" example:"true"`
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
	Data    T      `json:"data"`
}

type NotFoundErrorResponse struct {
	Ok      bool   `json:"ok" example:"false"`
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"not_found_error"`
}

type UnAuthorizedErrorResponse struct {
	Ok      bool   `json:"ok" example:"false"`
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"unauthorized_error"`
}

type InternalServerErrorResponse struct {
	Ok      bool   `json:"ok" example:"false"`
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal_server_error"`
}

type EntityErrorResponse struct {
	Ok      bool   `json:"ok" example:"false"`
	Code    int    `json:"code" example:"422"`
	Message string `json:"message" example:"entity_error"`
}
