package request_handler

import (
	"net/http"
)

func DisableAccessControl(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		inner.ServeHTTP(w, r)
	})
}
