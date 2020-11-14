package utils

// ResponseError ...
type ResponseError struct {
	message string
}

// NewResponseError ...
func NewResponseError(errorMessage string) ResponseError {
	return ResponseError{
		message: errorMessage,
	}
}

func (a ResponseError) Error() string {
	return a.message
}
