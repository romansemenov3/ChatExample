package request_handler

import (
	"encoding/json"
	"log"
	"model/service_error"
	"model/dto"
	"net/http"
)

func ErrorAdvice(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverIfNecessary(w, r)
		inner.ServeHTTP(w, r)
	})
}

func recoverIfNecessary(w http.ResponseWriter, r *http.Request) {
	if rec := recover(); rec != nil {
		handleError(rec, w, r)
	}
}

func handleError(rec interface{}, w http.ResponseWriter, r *http.Request) {
	err, isError := rec.(error)
	if !isError {
		log.Printf("Recovered %s", rec)
		err = service_error.WrapString("Unexpected error type")
	}

	serviceError, isServiceError := err.(service_error.ServiceError)
	if !isServiceError {
		serviceError = service_error.WrapError(err)
	}

	log.Printf(
		"%s (%s) %s",
		serviceError.Code(),
		serviceError.Title(),
		serviceError.Error(),
	)
	log.Print(err)

	body, _ := json.Marshal(dto.ErrorEntryDTO{
		Code:    serviceError.Code(),
		Title:   serviceError.Title(),
		Message: serviceError.Error(),
	})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(serviceError.Status())
	w.Write(body)
}