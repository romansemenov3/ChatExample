package index_controller

import (
	"api/controller"
	"fmt"
	"net/http"
)

const prefix = "/"

var RouteGet = controller.Route{
	Name:        "Index",
	Method:      "GET",
	Pattern:     prefix,
	HandlerFunc: index,
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chat is online, API at " + controller.Prefix)
}