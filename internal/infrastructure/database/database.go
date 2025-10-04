package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

type DatabaseInterface interface {
	GetInstance() *sql.DB
}

func CreateDatabaseInstance(db DatabaseInterface) DatabaseInterface {
	return db
}

type MySQL struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

func (m *MySQL) GetInstance() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", m.Username, m.Password, m.Host, m.Port, m.Name)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	})

	return db
}
