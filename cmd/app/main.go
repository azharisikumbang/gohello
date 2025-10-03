package main

import (
	"github.com/azharisikumbang/gohello/internal/infrastructure/app"
	database "github.com/azharisikumbang/gohello/internal/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.GetInstance()
	defer db.Close()

	app.Run()
}
