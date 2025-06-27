package models

import "time"

type User struct {
	ID        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"john_doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Password  string    `json:"-" swaggerignore:"true"`
	Role      string    `json:"role" example:"student"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T12:00:00Z"`
}
