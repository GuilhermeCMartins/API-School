package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"nome" validate:"nonzero"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
