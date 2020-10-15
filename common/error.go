package common

type ServerError struct {
	Message string `json:"message"`
}

func NewServerError(message string) *ServerErrorResponseBody {
	return &ServerErrorResponseBody{
		Error: &ServerError{
			Message: message,
		},
	}
}

type ServerErrorResponseBody struct {
	Error *ServerError `json:"error"`
}
