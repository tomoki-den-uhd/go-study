package models

import (
	"time"
)

// Attendance 出席テーブル
type Attendance struct {
	AttendanceID  int       `json:"attendance_id"`
	StudentUserID int       `json:"student_user_id"`
	CourseID      int       `json:"course_id"`
	Status        string    `json:"status"` // enumなのでstringやintどちらかに
	AttendedAt    time.Time `json:"attended_at"`
	IsDeleted     bool      `json:"is_deleted"`
} 