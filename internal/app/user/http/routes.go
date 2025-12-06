package http

import (
	contract "github.com/azharisikumbang/gohello/internal"
	"github.com/azharisikumbang/gohello/internal/app/user/domain"
	"github.com/azharisikumbang/gohello/internal/app/user/infra"
)

func Routes(r contract.RouterInterface, db contract.DatabaseInterface, cfg contract.AppConfig) {

	repo := infra.NewMySQLUserRepository(db.GetInstance())
	serv := domain.NewUserService(repo, cfg)
	h := NewUserHandler(serv)

	r.Post("/login", h.PostLoginHandler, nil)

	r.Get("/users", h.GetUsersHandler, Middlewares())
	r.Post("/users", h.PostUserHandler, nil)
}

func Middlewares() []contract.MiddlewareInterface {
	var middlewares []contract.MiddlewareInterface

	middlewares = append(middlewares, NewLogMiddlware())

	return middlewares
}
