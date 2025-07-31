package models

import (
	"time"

	"github.com/jackc/pgx/v5"
)

// User ユーザーの構造体
type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` // "student", "teacher", "parent" など
}

// Grade 成績の構造体
type Grade struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CourseID   int       `json:"course_id"`
	TotalScore float64   `json:"total_score"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDeleted  bool      `json:"is_deleted"`
}

// GradeDetail 成績詳細テーブル
type GradeDetail struct {
	GradeDetailID  int    `json:"grade_detail_id"`
	GradeID        int    `json:"grade_id"`
	TestQuestionID int    `json:"test_question_id"`
	StudentAnswer  string `json:"student_answer"`
	IsCorrect      bool   `json:"is_correct"`
	Score          int    `json:"score"`
}

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

// Attendance 出席テーブル
type Attendance struct {
	AttendanceID  int       `json:"attendance_id"`
	StudentUserID int       `json:"student_user_id"`
	CourseID      int       `json:"course_id"`
	Status        string    `json:"status"` // enumなのでstringやintどちらかに
	AttendedAt    time.Time `json:"attended_at"`
	IsDeleted     bool      `json:"is_deleted"`
}

// Question 質問テーブル
type Question struct {
	QuestionID        int       `json:"question_id"`
	StudentUserID     int       `json:"student_user_id"`
	SubjectID         int       `json:"subject_id"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	IsUnderstand      bool      `json:"is_understand"`
	IsPrivate         bool      `json:"is_private"`
	RelatedQuestionID *int      `json:"related_question_id"` // NULL許容
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsDeleted         bool      `json:"is_deleted"`
}

// Answer 回答テーブル
type Answer struct {
	AnswerID      int       `json:"answer_id"`
	QuestionID    int       `json:"question_id"`
	TeacherUserID int       `json:"teacher_user_id"`
	Content       string    `json:"content"`
	IsDraft       bool      `json:"is_draft"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsDeleted     bool      `json:"is_deleted"`
}

// Notification 通知テーブル
type Notification struct {
	NotificationID int       `json:"notification_id"`
	SenderUserID   int       `json:"sender_user_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	TargetRole     string    `json:"target_role"` // enum扱い
	SentAt         time.Time `json:"sent_at"`
	IsDeleted      bool      `json:"is_deleted"`
}

// ReadReceipt 既読管理テーブル
type ReadReceipt struct {
	ReceiptID      int       `json:"receipt_id"`
	NotificationID int       `json:"notification_id"`
	UserID         int       `json:"user_id"`
	ReadAt         time.Time `json:"read_at"`
	IsRead         bool      `json:"is_read"`
	IsDeleted      bool      `json:"is_deleted"`
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