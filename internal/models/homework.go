package models

import "time"

type Homework struct {
	ID          int       `json:"id"`
	LessonID    int       `json:"lesson_id"`
	StudentID   int       `json:"student_id"`
	Description string    `json:"description"`
	SubmittedAt time.Time `json:"submitted_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
