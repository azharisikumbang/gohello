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

	req := request.NewCreateRegistrationReq(r)
	h.service.RegisterNewAccount(req)

}

func (h *UserHandler) PostLoginHandler(w http.ResponseWriter, r *http.Request) {

	req := request.NewLoginRequest(r)

	if !h.service.AuthenticateUser(req.Username, req.Password) {
		respData := helper.NewStdReponse(nil, []error{errors.New("invalid credentials")})

		helper.ToJson(respData, w, http.StatusUnauthorized)
		return
	}

	token := h.service.CreateLoginToken(req.Username)

	fmt.Println(token)

	respData := helper.NewStdReponse(response.NewValidLoginReponse(token), nil)
	helper.ToJson(respData, w, http.StatusOK)
}
