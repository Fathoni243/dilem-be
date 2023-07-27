package model

import (
	"time"
	"web-desa/request"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(100)"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	ReadAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByUsername(username string) (*User, error) 
	Update(user *User) (*User, error)
	Delete(user *User) error
}

type UserService interface {
	Register(user *request.UserRequest) (*User, error)
	GetAllUser() ([]*User, error)
	GetByUsername(username string) (*User, error)
	EditUser(id uint, user *request.UserRequest) (*User, error)
	DestroyUser(id uint) error
}