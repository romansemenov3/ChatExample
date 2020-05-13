package message_repository

import (
	"context"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"model/entity"
	"service/repository"
	"time"
)

const asc = 1
const desc = -1

var collection *mongo.Collection

func init() {
	collection = repository.Database.Collection("messages")
}

func Create(message entity.MessageEntity) entity.MessageEntity {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	message.Id = id.String()
	message.Date = time.Now()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	body, err := bson.Marshal(message)
	if err != nil {
		panic(err)
	}
	collection.InsertOne(ctx, body)
	return message
}

func Find(chatId *string, limit int64, offset int64) []entity.MessageEntity {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	options := options.FindOptions{}
	options.SetLimit(limit)
	options.SetSkip(offset)
	options.SetSort(bson.M{"date": desc})

	filter := bson.M{}
	if chatId != nil {
		filter["chatId"] = *chatId
	}

	cursor, err := collection.Find(ctx, filter, &options)
	if err != nil {
		panic(err)
	}

	messages := make([]entity.MessageEntity, 0)
	for cursor.Next(ctx) {
		message := entity.MessageEntity{}
		err := cursor.Decode(&message)
		if err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}
	return messages
}
