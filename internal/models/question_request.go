package models

import (
	"time"
)

// QuestionCreateRequest 質問作成リクエストの構造体
type QuestionCreateRequest struct {
	SubjectID         int    `json:"subject_id" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Content           string `json:"content" validate:"required"`
	IsPrivate         bool   `json:"is_private"`
	RelatedQuestionID *int   `json:"related_question_id"`
}

// QuestionUpdateRequest 質問更新リクエストの構造体
type QuestionUpdateRequest struct {
	Title             string `json:"title" validate:"required"`
	Content           string `json:"content" validate:"required"`
	IsUnderstand      bool   `json:"is_understand"`
	IsPrivate         bool   `json:"is_private"`
	RelatedQuestionID *int   `json:"related_question_id"`
}

// QuestionListRequest 質問一覧取得リクエストの構造体
type QuestionListRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	SubjectID *int   `json:"subject_id"`
	IsPrivate *bool  `json:"is_private"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

// QuestionResponse 質問レスポンスの構造体
type QuestionResponse struct {
	QuestionID        int       `json:"question_id"`
	StudentUserID     int       `json:"student_user_id"`
	StudentName       string    `json:"student_name"`
	SubjectID         int       `json:"subject_id"`
	SubjectName       string    `json:"subject_name"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	IsUnderstand      bool      `json:"is_understand"`
	IsPrivate         bool      `json:"is_private"`
	RelatedQuestionID *int      `json:"related_question_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AnswerCount       int       `json:"answer_count"`
}

// QuestionListResponse 質問一覧レスポンスの構造体
type QuestionListResponse struct {
	Count     int               `json:"count"`
	Questions []QuestionResponse `json:"questions"`
}

// AnswerCreateRequest 回答作成リクエストの構造体
type AnswerCreateRequest struct {
	QuestionID int    `json:"question_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	IsDraft    bool   `json:"is_draft"`
}

// AnswerResponse 回答レスポンスの構造体
type AnswerResponse struct {
	AnswerID      int       `json:"answer_id"`
	QuestionID    int       `json:"question_id"`
	TeacherUserID int       `json:"teacher_user_id"`
	TeacherName   string    `json:"teacher_name"`
	Content       string    `json:"content"`
	IsDraft       bool      `json:"is_draft"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
} 