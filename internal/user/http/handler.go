package http

import (
	"net/http"

	"github.com/azharisikumbang/gohello/internal/user/domain"
	request "github.com/azharisikumbang/gohello/internal/user/http/requests"
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
