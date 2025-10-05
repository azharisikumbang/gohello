package domain

type User struct {
	Id       int
	Username string
}

type UserRepository interface {
	All() ([]User, error)
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
