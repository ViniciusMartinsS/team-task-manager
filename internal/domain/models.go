package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

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

type Role struct {
	gorm.Model
	Name string
}

type Task struct {
	gorm.Model
	Name      string
	Summary   string
	UserId    int
	Performed *time.Time
}

type LoginPayload struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	Code        int    `json:"status"`
	AccessToken string `json:"accessToken,omitempty"`
	Message     string `json:"message,omitempty"`
}

type Claims struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type TaskPayload struct {
	Name      string `json:",omitempty"`
	Summary   string `json:",omitempty"`
	Performed string `json:",omitempty"`
}

type TaskResponse struct {
	Code    int                   `json:"status"`
	Result  []TaskResponseContent `json:"result,omitempty"`
	Message string                `json:"message,omitempty"`
}

type TaskResponseContent struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Performed string `json:"Performed,omitempty"`
}

type HandleTaskRequest struct {
	Body   []byte
	UserId int
	TaskId string
}
