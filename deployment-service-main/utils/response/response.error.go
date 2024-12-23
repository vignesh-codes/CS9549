package response

import (
	"net/http"
)

type Type string

type Error struct {
	Type       Type   `json:"error_code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

const (
	ErrInvalidRoute          Type = "INVALID_ROUTE"
	ErrEmptyAuthorization    Type = "EMPTY_AUTHORIZATION"
	ErrUnauthorized          Type = "UNAUTHORIZED"
	ErrUserNotFound          Type = "USER_NOT_FOUND"
	ErrInvalidAppCredentials Type = "INVALID_APP_CREDENTIALS"
	ErrFailedToProcess       Type = "FAILED_TO_PROCESS"
	ErrRateLimitExceed       Type = "TOO_MANY_REQUESTS"
	ErrInvalidData           Type = "INVALID_DATA"
	ErrEmptyBody             Type = "EMPTY_BODY"
	ErrEmptyParam            Type = "EMPTY_PARAM"
	ErrValidationError       Type = "VALIDATION_ERROR"
	ErrInvalidJwtSignature   Type = "JWT_SIGNATURE_VERIFICATION_FAILED"
	ErrBadRequest            Type = "BAD_REQUEST"
	ErrUnsupportedMediaType  Type = "UNSUPPORTED_MEDIA_TYPE"
	ErrItemNotFound          Type = "ITEM_NOT_FOUND"
	ErrExternalServiceDown   Type = "EXTERNAL_SERVICE_DOWN"
)

func (e *Error) Status() int {
	switch e.Type {

	case ErrEmptyAuthorization:
		return http.StatusUnauthorized

	case ErrInvalidAppCredentials:
		return http.StatusUnauthorized

	case ErrUnauthorized:
		return http.StatusUnauthorized

	case ErrInvalidJwtSignature:
		return http.StatusUnauthorized

	case ErrBadRequest:
		return http.StatusBadRequest

	case ErrUserNotFound:
		return http.StatusNotFound

	case ErrFailedToProcess:
		return http.StatusInternalServerError

	case ErrInvalidData:
		return http.StatusBadRequest

	case ErrEmptyBody:
		return http.StatusBadRequest

	case ErrValidationError:
		return http.StatusBadRequest

	case ErrInvalidRoute:
		return http.StatusNotFound

	case ErrItemNotFound:
		return http.StatusBadRequest

	case ErrRateLimitExceed:
		return http.StatusTooManyRequests

	default:
		return http.StatusInternalServerError
	}
}

func ValidationError(errorType Type, message string) *Error {
	return &Error{
		Type:       errorType,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func RateLimitExceedError(errorType Type, message string) *Error {
	return &Error{
		Type:       errorType,
		Message:    message,
		StatusCode: http.StatusTooManyRequests,
	}
}

func InternalServerError(event string, action string, message error) *Error {
	return &Error{
		Type:       ErrFailedToProcess,
		Message:    message.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func BadRequest(message string) *Error {
	return &Error{
		Type:       ErrBadRequest,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func RouteNotFound() *Error {
	return &Error{
		Type:       ErrInvalidRoute,
		Message:    "Invalid route.",
		StatusCode: http.StatusNotFound,
	}
}

func InvalidAppCredentials(message string) *Error {
	return &Error{
		Type:       ErrInvalidAppCredentials,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func UnAuthorized(message string) *Error {
	return &Error{
		Type:       ErrUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func ItemNotFound(message string) *Error {
	return &Error{
		Type:       ErrItemNotFound,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func ExternalServiceDown(action string, message error) *Error {
	return &Error{
		Type:       ErrExternalServiceDown,
		Message:    message.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func UnsupportedMediaType(message string) *Error {
	return &Error{
		Type:       ErrUnsupportedMediaType,
		Message:    message,
		StatusCode: http.StatusUnsupportedMediaType,
	}
}
