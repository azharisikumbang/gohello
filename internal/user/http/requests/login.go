package request

import (
	"encoding/json"
	"net/http"
)

type LoginnRequest struct {
	Username string
	Password string
}

func NewLoginRequest(r *http.Request) *LoginnRequest {
	var data LoginnRequest
	json.NewDecoder(r.Body).Decode(&data)

	return &data
}
