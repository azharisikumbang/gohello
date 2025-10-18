package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegistrationRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_confirmation"`
}

func NewRegistrationRequest(r *http.Request) *RegistrationRequest {
	var data RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &data
}

func (r *RegistrationRequest) Validate() []map[string]interface{} {
	var e []map[string]interface{} = nil

	if r.Name == "" {
		e = append(e, map[string]interface{}{"name": "field name is required."})
	}

	if r.Username == "" {
		e = append(e, map[string]interface{}{"username": "field username is required."})
	}

	if r.Password == "" {
		e = append(e, map[string]interface{}{"password": "field password is required."})
	}

	if r.PasswordConf != r.Password {
		e = append(e, map[string]interface{}{"password_confirmation": "Password confirmation not match."})
	}

	return e
}
