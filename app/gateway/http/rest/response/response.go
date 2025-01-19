package response

import (
	"net/http"
)

type Response struct {
	StatusCode int
	Body       any
}

func Ok(body any) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

func Created(body any) *Response {
	return &Response{
		StatusCode: http.StatusCreated,
		Body:       body,
	}
}

func NoContent() *Response {
	return &Response{
		StatusCode: http.StatusNoContent,
	}
}

func BadRequest(err error) *Response {
	return &Response{
		StatusCode: http.StatusBadRequest,
		Body: Error{
			Code:    "app:bad_request",
			Message: err.Error(),
		},
	}
}

func InternalServerErr(err error) *Response {
	return &Response{
		StatusCode: http.StatusInternalServerError,
		Body: Error{
			Code:    "app:internal_server_error",
			Message: err.Error(),
		},
	}
}

func AppError(err error) *Response {
	appErr := ToAppErr(err)

	return &Response{
		StatusCode: appErr.StatusCode,
		Body:       appErr.Error,
	}
}
