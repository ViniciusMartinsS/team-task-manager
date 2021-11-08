package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Email    string
	Password string
	Age      int
	RoleID   int
	Role     Role
	Task     []Task
}
