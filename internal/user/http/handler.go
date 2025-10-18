package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/azharisikumbang/gohello/internal/user/domain"
	request "github.com/azharisikumbang/gohello/internal/user/http/requests"
	response "github.com/azharisikumbang/gohello/internal/user/http/responses"
	"github.com/azharisikumbang/gohello/pkg/helper"
)

type UserHandler struct {
	service *domain.UserService
}

func NewUserHandler(s *domain.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.Repo.All()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respData := helper.NewStdReponse(users, nil)
	helper.ToJson(respData, w, http.StatusOK)
}

func (h *UserHandler) PostUserHandler(w http.ResponseWriter, r *http.Request) {

	req := request.NewRegistrationRequest(r)
	err := req.Validate()
	if err != nil {
		fmt.Println(err)
		// helper.NewErrorJsonReponse(w, err, http.StatusBadRequest)
		return
	}

	err2 := h.service.RegisterNewAccount(req)
	if err2 != nil {
		helper.NewErrorJsonReponse(w, []error{err2}, http.StatusBadGateway)
		return
	}

	helper.NewOkJsonReponse(w, "user created.", http.StatusOK)
}

func (h *UserHandler) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	req := request.NewLoginRequest(r)

	if !h.service.AuthenticateUser(req.Username, req.Password) {
		helper.NewErrorJsonReponse(w, []error{errors.New("invalid credentials")}, http.StatusUnauthorized)
		return
	}

	token, err := h.service.CreateLoginToken(req.Username)
	if err != nil || token == "" {
		helper.NewErrorJsonReponse(w, []error{errors.New("invalid credentials")}, http.StatusUnauthorized)
		return
	}

	helper.NewOkJsonReponse(w, response.NewValidLoginReponse(token), http.StatusOK)
}
