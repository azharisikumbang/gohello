package app

import (
	"database/sql"
	"log"

	core "github.com/azharisikumbang/gohello/internal"
)

type DatabaseServer struct {
	cfg core.DBConfig
}

func NewDatabaseServer(cfg core.DBConfig) core.DatabaseInterface {
	return &DatabaseServer{
		cfg: cfg,
	}
}

func (s *DatabaseServer) GetInstance() *sql.DB {
	switch s.cfg.Driver {
	case "mysql":
		return NewMySQL(s.cfg).GetInstance()

	default:
		log.Fatal("Failed to get database instance.")
	}

	return nil
}
