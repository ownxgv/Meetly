package users

type Repo interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
}

type Service interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

type Context interface {
	JSON(code int, obj interface{}) error
	BindJSON(obj interface{}) error
}

type Handler interface {
	GetAllUsers(c Context)
	CreateUser(c Context)
	GetUserByID(c Context)
	UpdateUser(c Context)
	DeleteUser(c Context)
}
