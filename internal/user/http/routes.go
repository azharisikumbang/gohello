package http

import (
	"database/sql"

	"github.com/azharisikumbang/gohello/internal/user/domain"
	"github.com/azharisikumbang/gohello/internal/user/infra"
	"github.com/azharisikumbang/gohello/pkg/core"
)

func Routes(r core.RouterInterface, db *sql.DB, cfg core.AppConfig) {

	repo := infra.NewMySQLUserRepository(db)
	serv := domain.NewUserService(repo, cfg)
	h := NewUserHandler(serv)

	r.Post("/login", h.PostLoginHandler, nil)

	r.Get("/users", h.GetUsersHandler, Middlewares())
	r.Post("/users", h.GetUsersHandler, Middlewares())
}

func Middlewares() []core.MiddlewareInterface {
	var middlewares []core.MiddlewareInterface

	middlewares = append(middlewares, NewLogMiddlware())

	return middlewares
}
