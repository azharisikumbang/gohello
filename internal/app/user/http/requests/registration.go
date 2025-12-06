package request

import (
	"encoding/json"
	"net/http"

	"github.com/azharisikumbang/gohello/pkg/dto"
	"github.com/azharisikumbang/gohello/pkg/helper"
)

type RegistrationRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_confirmation"`
}

func NewRegistrationRequest(r *http.Request) (*RegistrationRequest, error) {
	var data RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RegistrationRequest) Validate() []dto.FormError {
	var e []dto.FormError

	if r.Name == "" {
		e = append(e, helper.NewFormError("name", "field name is required"))
	}

	if r.Username == "" {
		e = append(e, helper.NewFormError("username", "field username is required"))
	}

	if r.Password == "" {
		e = append(e, helper.NewFormError("password", "field password is required"))
	}

	if r.PasswordConf != r.Password {
		e = append(e, helper.NewFormError("password_confirmation", "password and password confirmation is not match"))
	}

	return e
}
