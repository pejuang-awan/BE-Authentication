package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Role     uint8  `gorm:"column:role;not null"`
	GameType uint8  `gorm:"column:game_type;not null"`
}

func (impl *User) TableName() string {
	return "users"
}
