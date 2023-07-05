package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	RoleID   uint   `gorm:"not null"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}
