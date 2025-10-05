package user_managememnt

type User struct {
	Id       int
	Username string
}

type IUserRepository interface {
	All() ([]User, error)
}

type UserService struct {
	Repo IUserRepository
}

func NewUserService(r IUserRepository) *UserService {
	return &UserService{
		Repo: r,
	}
}

func (s *UserService) All() ([]User, error) {
	return s.Repo.All()
}
