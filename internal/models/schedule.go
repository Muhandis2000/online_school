package models

import "time"

type Schedule struct {
	ID        int       `json:"id"`
	LessonID  int       `json:"lesson_id"`
	Group     string    `json:"group"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
