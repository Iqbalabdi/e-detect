package response

type ApiResponse struct {
	Status string `json:"status"`
	Data   interface{}
}
