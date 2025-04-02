package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"` // "admin" or "user"
}

