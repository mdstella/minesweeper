package errors

import (
	"fmt"
	"net/http"
)

// Response error to show an error like this:
// {
//     "error": {
//         "message": "Wrong column value should be between 0 and 9",
//         "type": "BAD_PARAM_ERROR",
//         "code": 400
//     }
// }
type ErrorResponse struct {
	Error CoreError `json:"error,omitempty"`
}

// Core error, all the business errors will be from this struct type
type CoreError struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (e CoreError) Error() string {
	return fmt.Sprintf("Message:'%s', Type:%s, Code:%d", e.Message, e.Type, e.Code)
}

// NewBadParamError creates a new bad param error on the core
func NewBadParamError(message string) *CoreError {
	return &CoreError{Message: message, Code: http.StatusBadRequest, Type: "BAD_PARAM_ERROR"}
}

// GetCoreError retrieve the pointer to be able to use when we do the error http encoding in main file
func GetCoreError(err error) *CoreError {
	if err == nil {
		return nil
	}
	if nerr2, ok := err.(*CoreError); ok {
		return nerr2
	}
	if nerr2, ok := err.(CoreError); ok {
		return &nerr2
	}
	return nil
}
