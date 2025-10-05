package user_managememnt

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from user page")
}
