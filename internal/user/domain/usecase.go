package domain

import (
	"log"

	request "github.com/azharisikumbang/gohello/internal/user/http/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	All() ([]User, error)
	FindByUsername(u string) (*User, error)
	UpdatePassword(u string, p string) bool
}

type UserService struct {
	Repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{
		Repo: r,
	}
}

func (s *UserService) All() ([]User, error) {
	return s.Repo.All()
}

func (s *UserService) RegisterNewAccount(r *request.CreateRegistrationReq) {
	// create user account
	// create
}

func (s *UserService) AuthenticateUser(username string, password string) bool {
	user, err := s.Repo.FindByUsername(username)

	if err != nil || user == nil {
		return false
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Error : %s", err.Error())

		return false
	}

	return true
}

func (s *UserService) CreateLoginToken(username string) string {
	return "this is valid token"
}
