package infra

import (
	"database/sql"
	"errors"
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

func (r *MySQLUserRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User

	row := r.DB.QueryRow("SELECT id, username, password, user_group_id, created_at from users where username = ?", username)

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.UserGroupId, &user.CreatedAt)
	if err != nil {
		return nil, row.Err()
	}

	return &user, nil
}

func (r *MySQLUserRepository) Save(u *domain.User) error {
	q := "INSERT INTO users (username, password) VALUES (?, ?)"

	_, err := r.DB.Exec(q, u.Username, u.Password)

	if err != nil {
		return errors.New("server error. Failed to create data, please contact the administrator")
	}

	return nil
}

func (r *MySQLUserRepository) UpdatePassword(username string, password string) bool {
	q := string("update users set password = ? where username = ?")
	_, err := r.DB.Exec(q, password, username)

	return err != nil
}
