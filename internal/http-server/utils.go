package http_server

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("write response: %v", err)
	}
}

func DecodeJson(r *http.Request, res any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(res); err != nil {
		return err
	}
	return nil

}
