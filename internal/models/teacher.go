package models

import "time"

type Teacher struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FullName  string    `json:"full_name"`
	Subject   string    `json:"subject"`
	Approved  bool      `json:"approved"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
