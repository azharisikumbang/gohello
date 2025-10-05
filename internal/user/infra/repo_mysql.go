package infra

import (
	"database/sql"

	"github.com/azharisikumbang/gohello/internal/user/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) All() ([]domain.User, error) {
	return nil, nil
}
