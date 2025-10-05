package user

import (
	"github.com/azharisikumbang/gohello/internal/user/http"
	"github.com/azharisikumbang/gohello/pkg/core"
)

type UserFeature struct{}

func NewUserFeature() *UserFeature {
	return &UserFeature{}
}

func (f *UserFeature) Boot(a *core.Application) {
	http.Routes(a.Router, a.Db)
}
