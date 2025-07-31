package models

import (
	"time"

	"github.com/jackc/pgx/v5"
)

// TeacherTest テスト（教師作成）の構造体
type TeacherTest struct {
	TeacherTestID  int       `json:"teacher_test_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	DurationMinutes int      `json:"duration_minutes"`
	CourseID       int       `json:"course_id"`
	CreatedBy      int       `json:"created_by"`
	IsDraft        bool      `json:"is_draft"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ScheduledAt    time.Time `json:"scheduled_at"`
	IsDeleted      bool      `json:"is_deleted"`
}

// StudentTest 学生の受験テスト構造体
type StudentTest struct {
	StudentTestID  int       `json:"student_test_id"`
	TeacherTestID  int       `json:"teacher_test_id"`
	StudentUserID  int       `json:"student_user_id"`
	Score          int       `json:"score"`
	Comment        string    `json:"comment"`
	SubmittedAt    time.Time `json:"submitted_at"`
	IsDeleted      bool      `json:"is_deleted"`
}

// TestQuestion テスト問題構造体
type TestQuestion struct {
	TestQuestionID int    `json:"test_question_id"`
	TeacherTestID  int    `json:"teacher_test_id"`
	QuestionText   string `json:"question_text"`
	CorrectAnswer  string `json:"correct_answer"`
	Score          int    `json:"score"`
	IsDeleted      bool   `json:"is_deleted"`
}

// StudentTestAnswer 学生の問題回答構造体
type StudentTestAnswer struct {
	GradeDetailID  int    `json:"grade_detail_id"`
	StudentTestID  int    `json:"student_test_id"`
	TestQuestionID int    `json:"test_question_id"`
	StudentAnswer  string `json:"student_answer"`
	IsCorrect      bool   `json:"is_correct"`
	Score          int    `json:"score"`
	GradeType      string `json:"grade_type"`
	IsDeleted      bool   `json:"is_deleted"`
}

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
}

// TestRepository テストリポジトリの構造体
type TestRepository struct {
	DB *pgx.Conn
}