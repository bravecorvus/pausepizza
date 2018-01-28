package response

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
