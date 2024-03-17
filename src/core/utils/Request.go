package utils

import (
	"encoding/json"
	"io"
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

func ApplicationJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
		pr, pw := io.Pipe()
		tee := io.TeeReader(r.Body, pw)
		r.Body = pr

		go func() {
			body, _ := io.ReadAll(tee)
			defer pw.Close()
			log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, string(body))
		}()

		next.ServeHTTP(w, r)
	})
}
