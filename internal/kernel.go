package app

import (
	"github.com/azharisikumbang/gohello/internal/user"
	"github.com/azharisikumbang/gohello/pkg/core"
)

func NewApp() *core.Application {
	a := core.NewDefault()

	a.UseHTTPServer(CreateHTTPServer())
	a.UseDatabase(CreateDB(a.Config.DB))
	a.UseRouter(CreateRouter())

	// register all features here
	a.AddFeature(user.NewUserFeature())

	return a
}

func CreateDB(cfg core.DBConfig) core.DatabaseInterface {
	return core.NewDatabaseServer(cfg)
}

func CreateRouter() core.RouterInterface {
	return core.NewRouter()
}

func CreateHTTPServer() core.HTTPServerInterface {
	return core.NewHTTPServer()
}
