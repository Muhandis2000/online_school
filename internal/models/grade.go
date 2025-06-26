package models

import "time"

type Grade struct {
	ID          int       `json:"id"`
	StudentID   int       `json:"student_id"`
	LessonID    int       `json:"lesson_id"`
	Value       int       `json:"value"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
