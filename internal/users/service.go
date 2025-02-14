package users

type service struct {
	repo UserRepository // Было: Repository, стало: UserRepository
}

// NewService создаёт экземпляр UserService
func NewService(repo UserRepository) UserService { // Было: Repository, стало: UserRepository
	return &service{repo: repo}
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *service) UpdateUser(user *User) error {
	return s.repo.UpdateUser(user)
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
