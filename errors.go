package errors

import (
	"fmt"
)

type Error interface {
	GetMessage() string
	GetCode() int
	GetInternalReason() string
	GetHTTPStatus() int
	GetDetails() map[string]any
	Error() string
}

type ErrorType struct {
	message        string
	code           int
	internalReason string
	hTTPStatus     int
	details        map[string]any
}

func (e ErrorType) Error() string {
	return fmt.Sprintf("Error: [%d] %s : %s. details: %v", e.GetCode(), e.GetMessage(), e.GetInternalReason(), e.GetDetails())
}

func NewError(
	message string,
	code int,
	internalReason string,
	HTTPStatus int,
	details map[string]any,
) Error {
	return ErrorType{
		message,
		code,
		internalReason,
		HTTPStatus,
		details,
	}
}

func (e ErrorType) GetMessage() string {
	return e.message
}

func (e ErrorType) GetCode() int {
	return e.code
}

func (e ErrorType) GetInternalReason() string {
	return e.internalReason
}

func (e ErrorType) GetHTTPStatus() int {
	return e.hTTPStatus
}

func (e ErrorType) GetDetails() map[string]any {
	if e.details == nil {
		return make(map[string]any)
	}
	return e.details
}

var UndefinedError = NewError(
	"undefined error occurred",
	0,
	"undefined error occured",
	500,
	nil,
)
