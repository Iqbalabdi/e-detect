package response

type ApiResponse struct {
	Status  string      `json:"message"`
	Message interface{} `json:"data"`
}

type StatisticResponse struct {
	TotalReport int64       `json:"totalReport"`
	TotalBank   int64       `json:"totalBank"`
	TotalPhone  int64       `json:"totalPhone"`
	TotalCost   interface{} `json:"totalCost"`
}
