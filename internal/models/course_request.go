package models

import (
	"time"
)

// CreateCourseRequest 授業登録リクエストの構造体
type CreateCourseRequest struct {
	Title         string    `json:"title" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	SubjectID     int       `json:"subject_id" validate:"required"`
	ScheduledAt   time.Time `json:"scheduled_at" validate:"required"`
}

// CreateCourseResponse 授業登録レスポンスの構造体
type CreateCourseResponse struct {
	Status string                 `json:"status"`
	Info   map[string]interface{} `json:"info"`
	Data   CourseData             `json:"data"`
}

// CourseData 授業データの構造体
type CourseData struct {
	CourseID      int       `json:"course_id"`
	TeacherUserID int       `json:"teacher_user_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	SubjectID     int       `json:"subject_id"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	UpdatedAt     time.Time `json:"updated_at"`
} 