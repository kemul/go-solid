package domain

type User struct {
	ID    int64  `json:"user_id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetByID(id int64) (*User, error)
	Store(user *User) error
}

type UserUsecase interface {
	GetUserByID(id int64) (*User, error)
	CreateUser(user *User) error
}
