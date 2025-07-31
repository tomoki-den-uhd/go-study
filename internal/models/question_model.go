package models

import (
	"time"
)

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