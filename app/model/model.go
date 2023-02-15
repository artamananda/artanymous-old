package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Question string `json:"question"`
}

type ViewMessage struct {
	CreatedAt time.Time `json:"created_at"`
	Question  string    `json:"question"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Msg string `json:"msg"`
}
