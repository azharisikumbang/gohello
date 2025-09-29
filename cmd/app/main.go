package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/goapp?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		db.QueryRow("SELECT id, username, password, created_at from users where id = ?", 1).Scan(&id, &username, &password, &createdAt)

		fmt.Fprintf(w, "ID : %d, Username: %s, Password: %s, %s", id, username, password, createdAt.Format("2006-01-02 15:04:05"))

		fmt.Fprintf(w, "Its rendered")
	})

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

}
