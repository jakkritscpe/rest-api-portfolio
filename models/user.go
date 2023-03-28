package models

import "gorm.io/gorm"

type User struct {
	Username string
	Password string
	Fullname string
	Nickname string
	Avatar   string
	gorm.Model
}
