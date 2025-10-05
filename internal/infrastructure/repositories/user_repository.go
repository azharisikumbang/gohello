package repositories

import (
	"database/sql"

	user_managememnt "github.com/azharisikumbang/gohello/internal/domain/user_management"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) All() ([]user_managememnt.User, error) {
	return nil, nil
}
