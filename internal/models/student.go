package models

import "time"

type Student struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FullName  string    `json:"full_name"`
	Group     string    `json:"group"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
