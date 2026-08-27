package response

import "net/http"

// Response define a JSON structure for http response
type Response[T any] struct {
	Status int           `json:"status"`
	Data   T             `json:"data,omitempty"`
	Error  *ErrorMessage `json:"error,omitempty"`
}

// ErrorMessage struct
type ErrorMessage struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// Factory is a function to init Response struct
func Factory[T any]() *Response[T] {
	return &Response[T]{}
}

// WithData is a function to set data
func (res *Response[T]) WithData(dt T) *Response[T] {
	res.Data = dt
	return res
}

// ErrorMessage is a function to set error message
func (res *Response[T]) ErrorMessage(message string) *Response[T] {
	if res.Error == nil {
		res.Error = new(ErrorMessage)
	}
	res.Error.Message = message
	return res
}

// ErrorReason is a function to set error reason
func (res *Response[T]) ErrorReason(reason string) *Response[T] {
	if res.Error == nil {
		res.Error = new(ErrorMessage)
	}
	res.Error.Reason = reason
	return res
}

// ==== standard response ====

// Success define success response
func (res *Response[T]) Success() *Response[T] {
	res.Status = http.StatusOK
	return res
}

// Created will build error response with 201 http code
func (res *Response[T]) Created() *Response[T] {
	res.Status = http.StatusCreated

	return res
}

// GeneralError will build general error response with
func (res *Response[T]) GeneralError(statusCode int) *Response[T] {
	res.Status = statusCode

	// set error info
	res.Error = &ErrorMessage{
		Reason:  "",
		Message: "failed",
	}

	return res
}

// BadRequest will build error response with 400 http code
func (res *Response[T]) BadRequest() *Response[T] {
	res.Status = http.StatusBadRequest

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, there are some missing request",
		Reason:  "bad request",
	}

	return res
}

// Unauthorized will build error response with 401 http code
func (res *Response[T]) Unauthorized() *Response[T] {
	res.Status = http.StatusUnauthorized

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, you are not authorized",
		Reason:  "unauthorized",
	}

	return res
}

// Forbidden will build error response with 403 http code
func (res *Response[T]) Forbidden() *Response[T] {
	res.Status = http.StatusForbidden

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, you are forbidden to access this feature",
		Reason:  "forbidden",
	}

	return res
}

// NotFound will build error response with 404 http code
func (res *Response[T]) NotFound() *Response[T] {
	res.Status = http.StatusNotFound

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, the feature that you're looking for is not found",
		Reason:  "not found",
	}

	return res
}

// Conflict will build error response with 409 http code
func (res *Response[T]) Conflict() *Response[T] {
	res.Status = http.StatusConflict

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, there is a duplicate data",
		Reason:  "bad request",
	}

	return res
}

// EntityTooLarge will build error response with 413 http code
func (res *Response[T]) EntityTooLarge() *Response[T] {
	res.Status = http.StatusRequestEntityTooLarge

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, your file is too large",
		Reason:  "bad request",
	}

	return res
}

// UnprocessableEntity will build error response with 422 http code
func (res *Response[T]) UnprocessableEntity() *Response[T] {
	res.Status = http.StatusUnprocessableEntity

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, your request cannot processed",
		Reason:  "unprocessable",
	}

	return res
}

// InternalServerError will build error response with 500 http code
func (res *Response[T]) InternalServerError() *Response[T] {
	res.Status = http.StatusInternalServerError

	// set error info
	res.Error = &ErrorMessage{
		Message: "Sorry, there are some error",
		Reason:  "internal server error",
	}

	return res
}
