package core

import (
	"database/sql"
	"log"
)

type DatabaseServer struct {
	cfg DBConfig
}

func NewDatabaseServer(cfg DBConfig) DatabaseInterface {
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
