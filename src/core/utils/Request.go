package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseMessage(message string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Message: message,
		Data:    data,
	}
}

func HandleExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(NewResponseMessage(r.(string), nil))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
