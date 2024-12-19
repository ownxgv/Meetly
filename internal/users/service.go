package users

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) UpdateUser(user *User) error {
	return s.repo.UpdateUser(user)
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
