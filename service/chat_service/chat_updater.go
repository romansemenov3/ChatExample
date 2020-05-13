package chat_service

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"log"
	"model/dto"
	"net/http"
)

type chatConnections map[uuid.UUID]*websocket.Conn //connectionId => connection
type connections map[string]chatConnections        //chatId => connectionPool

var activeConnections connections
var upgrader websocket.Upgrader

func init() {
	activeConnections = connections{}

	upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func AttachListener(chatId string, w http.ResponseWriter, r *http.Request) {
	connectionId, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	if _, ok := activeConnections[chatId]; !ok {
		activeConnections[chatId] = chatConnections{}
	}

	activeConnections[chatId][connectionId] = connection
	defer delete(activeConnections[chatId], connectionId)

	for {
		connection.ReadMessage()
	}
}

func PostMessage(message dto.MessageDTO) {
	if _, ok := activeConnections[message.ChatId]; !ok {
		return
	}

	for _, connection := range activeConnections[message.ChatId] {
		err := connection.WriteJSON(message)
		if err != nil {
			log.Println("write:", err)
		}
	}
}
