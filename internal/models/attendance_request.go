package models

import (
	"time"
)

// AttendanceCreateRequest 出席登録リクエストの構造体
type AttendanceCreateRequest struct {
	StudentUserID int    `json:"student_user_id" validate:"required"`
	CourseID      int    `json:"course_id" validate:"required"`
	Status        string `json:"status" validate:"required,oneof=present absent late"`
}

// AttendanceUpdateRequest 出席更新リクエストの構造体
type AttendanceUpdateRequest struct {
	Status string `json:"status" validate:"required,oneof=present absent late"`
}

// AttendanceListRequest 出席一覧取得リクエストの構造体
type AttendanceListRequest struct {
	StudentUserID *int   `json:"student_user_id"`
	CourseID      *int   `json:"course_id"`
	Status        string `json:"status"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
}

// AttendanceResponse 出席レスポンスの構造体
type AttendanceResponse struct {
	AttendanceID  int       `json:"attendance_id"`
	StudentUserID int       `json:"student_user_id"`
	StudentName   string    `json:"student_name"`
	CourseID      int       `json:"course_id"`
	CourseTitle   string    `json:"course_title"`
	Status        string    `json:"status"`
	AttendedAt    time.Time `json:"attended_at"`
}

// AttendanceListResponse 出席一覧レスポンスの構造体
type AttendanceListResponse struct {
	Count       int                 `json:"count"`
	Attendances []AttendanceResponse `json:"attendances"`
}

// AttendanceSummaryResponse 出席サマリーレスポンスの構造体
type AttendanceSummaryResponse struct {
	StudentUserID int    `json:"student_user_id"`
	StudentName   string `json:"student_name"`
	CourseID      int    `json:"course_id"`
	CourseTitle   string `json:"course_title"`
	TotalDays     int    `json:"total_days"`
	PresentDays   int    `json:"present_days"`
	AbsentDays    int    `json:"absent_days"`
	LateDays      int    `json:"late_days"`
	AttendanceRate float64 `json:"attendance_rate"`
} 