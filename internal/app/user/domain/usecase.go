package domain

import (
	"errors"
	"log"
	"strconv"
	"time"

	core "github.com/azharisikumbang/gohello/internal"
	request "github.com/azharisikumbang/gohello/internal/app/user/http/requests"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	All() ([]User, error)
	FindByUsername(u string) (*User, error)
	UpdatePassword(u string, p string) bool
	Save(u *User) error
}

type UserService struct {
	Repo UserRepository
	Cfg  core.AppConfig
}

func NewUserService(r UserRepository, cfg core.AppConfig) *UserService {
	return &UserService{
		Repo: r,
		Cfg:  cfg,
	}
}

func (s *UserService) All() ([]User, error) {
	return s.Repo.All()
}

func (s *UserService) RegisterNewAccount(r *request.RegistrationRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to create user")
	}

	var user *User = &User{
		Username: r.Username,
		Password: string(hashed),
	}

	return s.Repo.Save(user)
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

func (s *UserService) CreateLoginToken(username string) (string, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	subject := user.Id

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(subject),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.Cfg.Key))
}
