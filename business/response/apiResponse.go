package response

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
