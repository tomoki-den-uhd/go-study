package models

import (
	"time"
)

// GradeDetailResponse 成績詳細取得用のレスポンス構造体
type GradeDetailResponse struct {
	Status string      `json:"status"`
	Data   GradeData   `json:"data"`
}

// GradeData 成績データの構造体
type GradeData struct {
	GradeSummary  GradeSummary   `json:"grade_summary"`
	GradeDetails  []GradeDetail  `json:"grade_details"`
}

// GradeSummary 成績概要の構造体
type GradeSummary struct {
	GradeID         int       `json:"grade_id"`
	StudentUserID   int       `json:"student_user_id"`
	TeacherTestID   int       `json:"teacher_test_id"`
	Score           int       `json:"score"`
	Comment         string    `json:"comment"`
	TestTitle       string    `json:"test_title"`
	SubjectName     string    `json:"subject_name"`
	ScheduledAt     time.Time `json:"scheduled_at"`
	DurationMinutes int       `json:"duration_minutes"`
}

// GradeDetail 成績詳細の構造体
type GradeDetail struct {
	TestQuestionID int    `json:"test_question_id"`
	QuestionText   string `json:"question_text"`
	StudentAnswer  string `json:"student_answer"`
	IsCorrect      bool   `json:"is_correct"`
	Score          int    `json:"score"`
}

// GradeDetailRequest 成績詳細取得用のリクエスト構造体
type GradeDetailRequest struct {
	GradeID string `json:"grade_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
} 