package models

import "time"

type Attendance struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	LessonID  int       `json:"lesson_id"`
	Present   bool      `json:"present"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
