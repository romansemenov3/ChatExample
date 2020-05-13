package dto

import "time"

type MessageDTO struct {
	Id       string    `json:"id"`
	ChatId   string    `json:"chatId"`
	AuthorId string    `json:"authorId"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
}
