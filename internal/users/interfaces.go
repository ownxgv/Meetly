package users

import "meetly/internal/context"

// UserRepository определяет методы для работы с базой данных
type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

// UserService определяет бизнес-логику
type UserService interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

// HTTPContext абстрагирует работу с HTTP-запросами
//type HTTPContext interface {
//	JSON(code int, obj interface{}) error
//	BindJSON(obj interface{}) error
//	Param(key string) string
//}

// UserHandler определяет интерфейс для HTTP-обработчиков
type UserHandler interface {
	GetAllUsers(c context.Context)
	CreateUser(c context.Context)
	GetUserByID(c context.Context)
	UpdateUser(c context.Context)
	DeleteUser(c context.Context)
}
