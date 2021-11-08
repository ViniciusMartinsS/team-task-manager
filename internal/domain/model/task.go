package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name      string
	Summary   string
	UserId    int
	Performed *time.Time
}

type HandleTaskRequest struct {
	Body   []byte
	UserId int
	TaskId string
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
