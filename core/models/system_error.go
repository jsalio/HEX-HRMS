package models

import "fmt"

type SystemErrorType string
type SystemErrorLevel string
type SystemErrorCode int

const (
	SystemErrorTypeInternal   SystemErrorType  = "internal"
	SystemErrorTypeValidation SystemErrorType  = "validation"
	SystemErrorLevelInfo      SystemErrorLevel = "info"
	SystemErrorLevelWarning   SystemErrorLevel = "warning"
	SystemErrorLevelError     SystemErrorLevel = "error"
)

const (
	SystemErrorCodeInternal   SystemErrorCode = 500
	SystemErrorCodeValidation SystemErrorCode = 400
	SystemErrorCodeMigration  SystemErrorCode = 404
	SystemErrorCodeNone       SystemErrorCode = 0
)

type SystemError struct {
	Code    SystemErrorCode
	Type    SystemErrorType
	Level   SystemErrorLevel
	Message string
	Details any
}

func NewSystemError(code SystemErrorCode, _type SystemErrorType, level SystemErrorLevel, message string, details any) *SystemError {
	return &SystemError{
		Code:    code,
		Type:    _type,
		Level:   level,
		Message: message,
		Details: details,
	}
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Level, e.Type, e.Message)
}
