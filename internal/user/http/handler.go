package http

import (
	"encoding/json"
	"net/http"

	"github.com/azharisikumbang/gohello/internal/user/domain"
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

	reponsesData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reponsesData)
}
