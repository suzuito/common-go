package cgin

import "github.com/gin-gonic/gin"

// HTTPError ...
type HTTPError struct {
	Code    int
	Message string
	Err     error
}

// NewHTTPError ...
func NewHTTPError(
	Code int,
	Message string,
	Err error,
) *HTTPError {
	return &HTTPError{
		Code:    Code,
		Message: Message,
		Err:     Err,
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

func (e *HTTPError) response() *errorResponse {
	return &errorResponse{
		Message: e.Message,
	}
}

// Abort ...
func Abort(ctx *gin.Context, err *HTTPError) {
	response := err.response()
	ctx.AbortWithStatusJSON(err.Code, &response)
}
