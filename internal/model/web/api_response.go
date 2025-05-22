package web

type ApiResponse struct {
	Ok      bool   `json:"ok" example:"true"`
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
	Data    any    `json:"data"`
}

type NotFoundResponse struct {
	Ok      bool   `json:"ok" example:"false"`
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"not_found_error"`
}
