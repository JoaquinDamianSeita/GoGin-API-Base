package responses_utils

type ApiErrorResponse struct {
	Error string `json:"error"`
}

type ApiMessageResponse struct {
	Message string `json:"message"`
}