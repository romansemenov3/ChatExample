package entity

import (
	"time"
)

type MessageEntity struct {
	Id       string    `bson:"_id"`
	ChatId   string    `bson:"chatId"`
	AuthorId string    `bson:"authorId"`
	Message  string    `bson:"message"`
	Date     time.Time `bson:"date"`
}
