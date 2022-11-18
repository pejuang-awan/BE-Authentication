package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
	Role     string `gorm:"column:role;not null"`
}

func (impl *User) TableName() string {
	return "users"
}
