package app

import "github.com/azharisikumbang/gohello/internal/infrastructure/database"

func NewDatabase(cfg DBConfig) DatabaseInterface {
	var db DatabaseInterface

	switch cfg.Driver {
	case "mysql":
		db = database.NewMySQL(
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.Name,
			cfg.Port,
		)
	}

	return db
}
