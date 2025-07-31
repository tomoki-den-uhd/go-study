package models

import (
	"time"
)

// Subject 教科テーブル
type Subject struct {
	SubjectID int       `json:"subject_id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Course コーステーブル
type Course struct {
	CourseID      int       `json:"course_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	TeacherUserID int       `json:"teacher_user_id"`
	SubjectID     int       `json:"subject_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	IsDeleted     bool      `json:"is_deleted"`
}

// CourseVideo コース動画テーブル
type CourseVideo struct {
	VideoID   int       `json:"video_id"`
	CourseID  int       `json:"course_id"`
	Filename  string    `json:"filename"`
	URL       string    `json:"url"`
	UploadedAt time.Time `json:"uploaded_at"`
	IsDeleted bool      `json:"is_deleted"`
} 