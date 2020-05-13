package chat_repository

import (
	"context"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"model/entity"
	"service/repository"
	"time"
)

var collection *mongo.Collection

func init() {
	collection = repository.Database.Collection("chats")
}

func Create(chat entity.ChatEntity) entity.ChatEntity {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	chat.Id = id.String()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	body, err := bson.Marshal(chat)
	if err != nil {
		panic(err)
	}
	collection.InsertOne(ctx, body)
	return chat
}

func GetById(id string) *entity.ChatEntity {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		panic(err)
	}

	if cursor.Next(ctx) {
		chat := new(entity.ChatEntity)
		err := cursor.Decode(&chat)
		if err != nil {
			panic(err)
		}
		return chat
	}
	return nil
}
