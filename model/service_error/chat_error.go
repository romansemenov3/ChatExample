package service_error

import "net/http"

type ChatNotFoundError struct {
	Id string
}

func (e ChatNotFoundError) Error() string {
	return "Chat with id " + e.Id + " not found"
}

func (e ChatNotFoundError) Title() string {
	return "Entity not found"
}

func (e ChatNotFoundError) Code() string {
	return prefix + "1001"
}

func (e ChatNotFoundError) Status() int {
	return http.StatusNotFound
}
