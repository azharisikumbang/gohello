package main

import (
	"github.com/azharisikumbang/gohello/internal/infrastructure/app"
)

func main() {
	app := app.NewDefault()
	app.Run()
}
