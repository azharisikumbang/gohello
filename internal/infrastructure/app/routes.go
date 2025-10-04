package app

import "github.com/azharisikumbang/gohello/internal/domain/user"

func (app *Application) LoadRoutes() {
	app.Get("/users", user.HomeHandler)
}
