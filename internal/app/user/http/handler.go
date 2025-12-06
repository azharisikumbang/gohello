package http

import (
	"net/http"

	"github.com/azharisikumbang/gohello/internal/app/user/domain"
	request "github.com/azharisikumbang/gohello/internal/app/user/http/requests"
	response "github.com/azharisikumbang/gohello/internal/app/user/http/responses"
	"github.com/azharisikumbang/gohello/pkg/dto"
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

	respData := helper.NewStdResponse(users, nil)
	helper.WriteJson(w, http.StatusOK, respData)
}

func (h *UserHandler) PostUserHandler(w http.ResponseWriter, r *http.Request) {

	req, err := request.NewRegistrationRequest(r)
	errs := req.Validate()
	if errs != nil {
		helper.NewBadRequestJsonResponse(w, errs)
		return
	}

	_, err = h.service.Repo.FindByUsername(req.Username)
	if err != nil {
		erv := helper.NewFormError("username", "Username already exists, try another.")
		helper.NewBadRequestJsonResponse(w, []dto.FormError{erv})
		return
	}

	err = h.service.RegisterNewAccount(req)
	if err != nil {
		helper.NewServerErrorJsonResponse(w, err)
		return
	}

	helper.NewOKJSONReponse(w, "user created.", http.StatusOK)
}

func (h *UserHandler) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewLoginRequest(r)

	if err != nil {
		helper.NewServerErrorJsonResponse(w, err)
	}

	if !h.service.AuthenticateUser(req.Username, req.Password) {
		erv := helper.NewFormError("login", "invalid credentials")
		helper.NewBadRequestJsonResponse(w, []dto.FormError{erv})
		return
	}

	token, err := h.service.CreateLoginToken(req.Username)
	if err != nil || token == "" {
		helper.NewServerErrorJsonResponse(w, err)
		return
	}

	helper.NewOKJSONReponse(w, response.NewValidLoginResponse(token), http.StatusOK)
}
