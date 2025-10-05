package main

import (
	"github.com/azharisikumbang/gohello/internal/features"
	"github.com/azharisikumbang/gohello/internal/infrastructure/app"
)

func main() {

	a := app.NewDefault()

	a.AddFeature(features.NewUserFeature())

	a.Run()
}
