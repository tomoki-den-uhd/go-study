package models

import (
	"time"
)

// TestListResponse 小テスト一覧表示用のレスポンス構造体
type TestListResponse struct {
	TeacherTestID   int       `json:"teacher_test_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DurationMinutes int       `json:"duration_minutes"`
	CourseTitle     string    `json:"course_title"`
	SubjectName     string    `json:"subject_name"`
	TeacherName     string    `json:"teacher_name"`
	ScheduledAt     time.Time `json:"scheduled_at"`
	IsDraft         bool      `json:"is_draft"`
	CreatedAt       time.Time `json:"created_at"`
	Comment         string    `json:"comment,omitempty"`
	Score           *int      `json:"score,omitempty"`
}

// TestListRequest 小テスト一覧取得用のリクエスト構造体
type TestListRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role" validate:"required,oneof=student teacher admin"`
}

// TestListResponseWrapper 小テスト一覧レスポンスのラッパー
type TestListResponseWrapper struct {
	Count int                `json:"count"`
	Tests []TestListResponse `json:"tests"`
} 