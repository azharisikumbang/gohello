package request

import "net/http"

type CreateRegistrationReq struct {
	Name         string
	Username     string
	Password     string
	PasswordConf string
}

func NewCreateRegistrationReq(r *http.Request) *CreateRegistrationReq {
	return &CreateRegistrationReq{
		Name:         r.FormValue("name"),
		Username:     r.FormValue("username"),
		Password:     r.FormValue("password"),
		PasswordConf: r.FormValue("password_confirmation"),
	}
}
