package response

type ErrorResponse struct {
	Message string `json:"message"`
	Status string `json:"status"`
}
