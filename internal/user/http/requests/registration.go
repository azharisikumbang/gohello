package request

import (
	"encoding/json"
	"fmt"
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

func NewRegistrationRequest(r *http.Request) *RegistrationRequest {
	var data RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &data
}

func (r *RegistrationRequest) Validate() []dto.ErrValue {
	var e []dto.ErrValue

	if r.Name == "" {
		e = append(e, helper.NewErrValue("name", "field name is required"))
	}

	if r.Username == "" {
		e = append(e, helper.NewErrValue("username", "field username is required"))
	}

	if r.Password == "" {
		e = append(e, helper.NewErrValue("password", "field password is required"))
	}

	if r.PasswordConf != r.Password {
		e = append(e, helper.NewErrValue("password_confirmation", "password and password confirmation is not match"))
	}

	return e
}
