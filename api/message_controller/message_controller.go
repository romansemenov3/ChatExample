package message_controller

import (
	"api/controller"
	"encoding/json"
	"model/dto"
	"net/http"
	"regexp"
	"service/message_service"
	"service/security_service"
	"strconv"
)

const prefix = controller.Prefix + "/message"
var rangePattern = regexp.MustCompile("(\\d+)-(\\d+)")

var RouteGet = controller.Route{
	Name:        "GetMessages",
	Method:      "GET",
	Pattern:     prefix,
	HandlerFunc: GetMessages,
}
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var chatId *string = nil
	chatIdParam, hasChatIdParam := r.URL.Query()["chatId"]
	if hasChatIdParam && len(chatIdParam[0]) > 0 {
		chatId = &chatIdParam[0]
	}

	var limit int64 = 10
	var offset int64 = 0
	rangeHeader, hasRangeHeader := r.Header["Range"]
	if hasRangeHeader && len(rangeHeader[0]) > 0 {
		matches := rangePattern.FindStringSubmatch(rangeHeader[0])
		start, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			panic("\"Range\" header format is invalid")
		}
		end, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			panic("\"Range\" header format is invalid")
		}

		offset = start - 1
		limit = end - start + 1
	}

	messages := message_service.Find(chatId, limit, offset)
	body, err := json.Marshal(messages)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

var RouteCreate = controller.Route{
	Name:        "CreateMessage",
	Method:      "POST",
	Pattern:     prefix,
	HandlerFunc: create,
}
func create(w http.ResponseWriter, r *http.Request) {
	var message dto.MessageDTO
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		panic(err)
	}

	authorId := security_service.GetUser(r)
	message = message_service.Create(message, authorId)
	body, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}