package models

import "time"

type Payment struct {
	ID          int       `json:"id"`
	StudentID   int       `json:"student_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
