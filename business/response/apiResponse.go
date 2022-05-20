package response

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type StatisticResponse struct {
	TotalReport int64       `json:"totalReport"`
	TotalBank   int64       `json:"totalBank"`
	TotalPhone  int64       `json:"totalPhone"`
	TotalCost   interface{} `json:"totalCost"`
}
