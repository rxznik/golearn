package response

type ResponseError struct {
	Error  string `json:"error" required:"true"`
	Reason string `json:"reason" required:"true"`
}

type Response struct {
	OK    any
	Error ResponseError
}
