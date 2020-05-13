package chat_service

import (
	"model/dto"
	"model/entity"
	"model/service_error"
	"service/chat_repository"
)

func Create(chat dto.ChatDTO) dto.ChatDTO {
	entity := entity.ChatEntity{
		Name: chat.Name,
	}
	entity = chat_repository.Create(entity)
	chat.Id = entity.Id
	return chat
}

func GetByIdOrThrow(id string) dto.ChatDTO {
	entity := chat_repository.GetById(id)
	if entity == nil {
		panic(service_error.ChatNotFoundError{Id: id})
	}
	return dto.ChatDTO{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
