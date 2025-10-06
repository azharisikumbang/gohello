package infra

import (
	"database/sql"
	"log"

	"github.com/azharisikumbang/gohello/internal/user/domain"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{
		DB: db,
	}
}

func (r *MySQLUserRepository) All() ([]domain.User, error) {
	var users []domain.User

	if r.DB == nil {
		panic("Dtabase null")
	}

	rows, err := r.DB.Query("SELECT id, username from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user = domain.User{}
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			log.Fatalf("Database Error : %s", err)
		}

		users = append(users, user)
	}

	return users, nil

}
