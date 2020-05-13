package dto

type ErrorEntryDTO struct {
	Code    string `json:"code,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}
