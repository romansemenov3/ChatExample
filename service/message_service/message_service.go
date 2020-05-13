package message_service

import (
	"common"
	"encoding/json"
	"fmt"
	"log"
	"model/dto"
	"model/entity"
	"service/chat_service"
	"service/message_repository"
	"time"

	"github.com/streadway/amqp"
)

type config struct {
	Amqp amqpProperties `yaml:"amqp"`
}

type amqpProperties struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var channel *amqp.Channel
var messages <-chan amqp.Delivery

func init() {
	cfg := config{}
	common.ReadConfig(&cfg)

	amqpURL := "amqp://" +
		cfg.Amqp.Username + ":" +
		cfg.Amqp.Password + "@" +
		cfg.Amqp.Host + ":" +
		cfg.Amqp.Port

	var conn *amqp.Connection
	var err error

	for {
		conn, err = amqp.Dial(amqpURL)
		if err != nil {
			fmt.Println("Could not connect to queue manager:", err)
			time.Sleep(time.Minute)
		} else {
			break
		}
	}

	channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	args := make(amqp.Table)
	_, err = channel.QueueDeclare("messages", true, false, false, false, args)
	if err != nil {
		log.Fatal(err)
	}

	messages, err = consumeMessages(channel)
	if err != nil {
		panic(err)
	}

	go func() {
		for delivery := range messages {
			messageDTO := dto.MessageDTO{}
			err := json.Unmarshal(delivery.Body, &messageDTO)
			if err != nil {
				log.Print(err)
			}
			chat_service.PostMessage(messageDTO)
			delivery.Ack(false)
		}
	}()
}

func publishMessage(message []byte, ch *amqp.Channel) error {
	return ch.Publish(
		"",         // exchange
		"messages", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{ContentType: "text/plain", Body: message},
	)
}

func consumeMessages(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		"messages", // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
}

func Create(message dto.MessageDTO, authorId string) dto.MessageDTO {
	chat_service.GetByIdOrThrow(message.ChatId)

	message.AuthorId = authorId
	entity := entity.MessageEntity{
		ChatId:   message.ChatId,
		AuthorId: message.AuthorId,
		Message:  message.Message,
	}
	entity = message_repository.Create(entity)
	message.Id = entity.Id
	message.Date = entity.Date

	body, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	publishMessage(body, channel)

	return message
}

func Find(chatId *string, limit int64, offset int64) []dto.MessageDTO {
	messages := make([]dto.MessageDTO, 0)
	for _, entity := range message_repository.Find(chatId, limit, offset) {
		message := dto.MessageDTO{
			Id:       entity.Id,
			ChatId:   entity.ChatId,
			AuthorId: entity.AuthorId,
			Message:  entity.Message,
			Date:     entity.Date,
		}
		messages = append(messages, message)
	}
	return messages
}
