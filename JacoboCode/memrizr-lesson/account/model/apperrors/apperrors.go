package apperrors

import (
	"fmt"
	"net/http"
)

// Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorType
const (
	Authorization        Type = "AUTHORIZATION"          // Authentication Failures
	BadRequest           Type = "BADREQUEST"             // Validation error / badinput
	Conflict             Type = "CONFLICT"               // ALready Exists (eg, create account with existent email) 409
	Internal             Type = "INTERNAL"               // Server (500) and fallback error
	NotFound             Type = "NOTFOUND"               // Fot not found finding resource
	PayloadTooLarge      Type = "PAYLOADTOOLARGE"        // for uploading tons of JSON, or an imageover the limit - 412
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"    // For long running handlers
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE" // for htt
)

// Error holds a custom error for the application
// which is helpful in returning a consistent
// error type/message from API endpoints
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies standard error interface
// we can return error form this package as
// a reguler old go _error_
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}

}

// Status check the runtime type
// of the error and return an http
// status code if the error is model.Error
func Status(err error) int {
	// var e *Error
	// if errors.As(err, &err) {
	// 	return e.Status()
	// }

	if e, ok := err.(*Error); ok {
		return e.Status()
	}

	return http.StatusInternalServerError
}

//ERROR FACTORIES

// NewAuthorization to create a 401
func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

// NewBadRequest to create 400 errors (validation, for example)
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflict to create an error for 409
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

// NewPayloadTooLarge to create an error for 413
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {

	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}

}

// NewServiceUnavailable to create an error for 503
func NewServiceUnavailable() *Error {
	return &Error{
		Type:    ServiceUnavailable,
		Message: fmt.Sprintf("Service unavailable or timed out"),
	}
}

// NewUnsupportedMediaType to create an error for 415
func NewUnsupportedMediaType(reason string) *Error {
	return &Error{
		Type:    UnsupportedMediaType,
		Message: reason,
	}
}
