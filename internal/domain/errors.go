package domain

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrCodeValidation     ErrorCode = "VALIDATION_ERROR"
	ErrCodeUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden      ErrorCode = "FORBIDDEN"
	ErrCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrCodeConflict       ErrorCode = "CONFLICT"
	ErrCodeRateLimited    ErrorCode = "RATE_LIMITED"
	ErrCodeInternal       ErrorCode = "INTERNAL_ERROR"
	ErrCodeTokenInvalid   ErrorCode = "TOKEN_INVALID"
	ErrCodeTokenExpired   ErrorCode = "TOKEN_EXPIRED"
	ErrCodeTokenReused    ErrorCode = "TOKEN_REUSED"
)

type AppError struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	RequestID string    `json:"request_id,omitempty"`
	Status    int       `json:"-"`
}

func (e *AppError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("%s: %s (request_id=%s)", e.Code, e.Message, e.RequestID)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) WithRequestID(requestID string) *AppError {
	cp := *e
	cp.RequestID = requestID
	return &cp
}

func NewAppError(code ErrorCode, message string, status int) *AppError {
	return &AppError{Code: code, Message: message, Status: status}
}

var (
	ErrValidation   = NewAppError(ErrCodeValidation, "validation failed", http.StatusBadRequest)
	ErrUnauthorized = NewAppError(ErrCodeUnauthorized, "unauthorized", http.StatusUnauthorized)
	ErrForbidden    = NewAppError(ErrCodeForbidden, "forbidden", http.StatusForbidden)
	ErrNotFound     = NewAppError(ErrCodeNotFound, "resource not found", http.StatusNotFound)
	ErrConflict     = NewAppError(ErrCodeConflict, "conflict", http.StatusConflict)
	ErrRateLimited  = NewAppError(ErrCodeRateLimited, "too many requests", http.StatusTooManyRequests)
	ErrInternal     = NewAppError(ErrCodeInternal, "internal server error", http.StatusInternalServerError)
	ErrTokenInvalid = NewAppError(ErrCodeTokenInvalid, "invalid token", http.StatusUnauthorized)
	ErrTokenExpired = NewAppError(ErrCodeTokenExpired, "token expired", http.StatusUnauthorized)
	ErrTokenReused  = NewAppError(ErrCodeTokenReused, "token reuse detected", http.StatusUnauthorized)
)
