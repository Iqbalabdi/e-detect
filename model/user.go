package model

import "gorm.io/gorm"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type User struct {
	*gorm.Model
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Phone      string `json:"phone" form:"phone"`
	Role       string `gorm:"default:user"`
	Reports    []Report
	Disclaimer []Disclaimer
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByQuery(user User) (User, error)
	Create(user User) (User, error)
	Update(id int, user User) (User, error)
}

type UserUseCase interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
	Validate(user *User) (bool, error)
	Login(user LoginRequest) (User, bool, error)
	Update(id int, user User) (User, error)
}
