package users

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *Repository) DeleteUser(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
