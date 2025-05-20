package web

type ApiResponse struct {
	Ok      bool   `json:"ok"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
