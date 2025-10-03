package app

import (
	"net/http"

	"github.com/azharisikumbang/gohello/internal/domain/user/handlers"
)

func LoadRoutes(h *http.ServeMux) {
	// domain/users
	h.HandleFunc("GET /users", handlers.HomeHandler)
}
