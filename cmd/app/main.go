package main

import (
	core "github.com/azharisikumbang/gohello/internal"
	"github.com/azharisikumbang/gohello/internal/app/user"
	"github.com/azharisikumbang/gohello/pkg/app"
)

func main() {
	a := app.NewDefault()

	a.UseHTTPServer(CreateHTTPServer())
	a.UseDatabase(CreateDB(a.GetConfig().DB))
	a.UseRouter(CreateRouter())

	// register all features here
	a.AddFeature(user.NewUserFeature())

	a.Run()
}

func CreateDB(cfg core.DBConfig) core.DatabaseInterface {
	return app.NewDatabaseServer(cfg)
}

func CreateRouter() core.RouterInterface {
	return app.NewRouter()
}

func CreateHTTPServer() core.HTTPServerInterface {
	return app.NewHTTPServer()
}
