package database

import (
	"database/sql"
	"sync"
	"time"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetInstance() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@(127.0.0.1:3306)/goapp?parseTime=true")
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
