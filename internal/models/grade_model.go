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