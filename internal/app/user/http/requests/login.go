package request

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string
	Password string
}

func NewLoginRequest(r *http.Request) (*LoginRequest, error) {
	var data LoginRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
