package chat_controller

import (
	"api/controller"
	"encoding/json"
	"model/dto"
	"net/http"
	"service/chat_service"
	"strings"
)

const prefix = controller.Prefix + "/chat"

var RouteGet = controller.Route{
	Name:        "GetChat",
	Method:      "GET",
	Pattern:     prefix + "/{id}",
	HandlerFunc: GetById,
}
func GetById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, prefix + "/")
	chat := chat_service.GetByIdOrThrow(id)
	body, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

var RouteCreate = controller.Route{
	Name:        "CreateChat",
	Method:      "POST",
	Pattern:     prefix,
	HandlerFunc: create,
}
func create(w http.ResponseWriter, r *http.Request) {
	var chat dto.ChatDTO
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		panic(err)
	}

	chat = chat_service.Create(chat)
	body, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

var RouteListen = controller.Route{
	Name:        "ListenChat",
	Method:      "GET",
	Pattern:     prefix + "/{id}/listen",
	HandlerFunc: listen,
}
func listen(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, prefix + "/"), "/listen")
	chat_service.AttachListener(id, w, r)
}

