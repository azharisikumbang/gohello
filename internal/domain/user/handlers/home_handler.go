package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/azharisikumbang/gohello/internal/infrastructure/database"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetUsers() ([]User, error) {
	var users []User
	var db *sql.DB = database.GetInstance()

	rows, err := db.Query("SELECT id, username from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
