package service_error

import "net/http"

type ServiceError interface {
	Code() string
	Title() string
	Error() string
	Status() int
}

const prefix = "CT-"

func WrapString(message string) ServiceError {
	return InternalError{Message: message}
}

func WrapError(e error) ServiceError {
	return WrapString(e.Error())
}

type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

func (e InternalError) Title() string {
	return "Unexpected error"
}

func (e InternalError) Code() string {
	return prefix + "0001"
}

func (e InternalError) Status() int {
	return http.StatusInternalServerError
}

type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "Invalid credentials"
}

func (e UnauthorizedError) Title() string {
	return "Unauthorized"
}

func (e UnauthorizedError) Code() string {
	return prefix + "0002"
}

func (e UnauthorizedError) Status() int {
	return http.StatusUnauthorized
}

type ForbiddenError struct {
}

func (e ForbiddenError) Error() string {
	return "Forbidden"
}

func (e ForbiddenError) Title() string {
	return "Forbidden"
}

func (e ForbiddenError) Code() string {
	return prefix + "0003"
}

func (e ForbiddenError) Status() int {
	return http.StatusForbidden
}
