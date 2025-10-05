package features

import (
	user_managememnt "github.com/azharisikumbang/gohello/internal/domain/user_management"
	"github.com/azharisikumbang/gohello/internal/infrastructure/app"
	"github.com/azharisikumbang/gohello/internal/infrastructure/repositories"
)

type UserFeature struct{}

func NewUserFeature() *UserFeature {
	return &UserFeature{}
}

func (f *UserFeature) Boot(a *app.Application) {
	f.routes(a)
}

func (f *UserFeature) routes(a *app.Application) {

	r := repositories.NewUserRepository(a.Db)
	s := user_managememnt.NewUserService(r)
	h := user_managememnt.NewUserHandler(s)

	a.Router.Get("/users", h.GetUsersHandler)
}
