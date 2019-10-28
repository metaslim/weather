package handler

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, payload json.Marshaler, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if payload == nil {
		msg := "Unexpected payload while parsing response"
		w.WriteHeader(statusCode)
		w.Write([]byte(msg))
		return
	}

	payloadBytes, err := payload.MarshalJSON()
	if err != nil {
		msg := "Unexpected error while parsing JSON response"
		w.WriteHeader(statusCode)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(payloadBytes)
}
