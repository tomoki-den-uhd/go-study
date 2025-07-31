package models

import (
	"time"
)

// Grade 成績の構造体
type Grade struct {
	ID             int       `json:"id"`
	StudentTestID  int       `json:"student_test_id"`
	StudentUserID  int       `json:"student_user_id"`
	CourseID       int       `json:"course_id"`
	TotalScore     float64   `json:"total_score"`
	Comment        string    `json:"comment"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
}

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

// GradeDetailTable 成績詳細テーブル（設計書対応）
type GradeDetailTable struct {
	GradeDetailID  int    `json:"grade_detail_id"`
	GradeID        int    `json:"grade_id"`
	TestQuestionID int    `json:"test_question_id"`
	StudentAnswer  string `json:"student_answer"`
	IsCorrect      bool   `json:"is_correct"`
	Score          int    `json:"score"`
}

// StudentTestAnswerTable 学生テスト回答テーブル（設計書対応）
type StudentTestAnswerTable struct {
	GradeDetailID  int    `json:"grade_detail_id"`
	StudentTestID  int    `json:"student_test_id"`
	TestQuestionID int    `json:"test_question_id"`
	StudentAnswer  string `json:"student_answer"`
	IsCorrect      bool   `json:"is_correct"`
	Score          int    `json:"score"`
	IsDeleted      bool   `json:"is_deleted"`
} 