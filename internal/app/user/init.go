package user

import (
	core "github.com/azharisikumbang/gohello/internal"
	"github.com/azharisikumbang/gohello/internal/app/user/http"
)

type Feature struct{}

func NewUserFeature() *Feature {
	return &Feature{}
}

func (f *Feature) Boot(a core.ApplicationInterface) {
	http.Routes(a.GetRouter(), a.GetDatabase(), a.GetConfig().App)
}
