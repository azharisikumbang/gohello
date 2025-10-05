package http

import (
	"fmt"
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
	fmt.Fprint(w, "Hello from user page 2")
}
