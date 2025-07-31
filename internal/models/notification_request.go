package models

import (
	"time"
)

// NotificationCreateRequest 通知作成リクエストの構造体
type NotificationCreateRequest struct {
	Title      string `json:"title" validate:"required"`
	Body       string `json:"body" validate:"required"`
	TargetRole string `json:"target_role" validate:"required,oneof=student teacher parent all"`
}

// NotificationListRequest 通知一覧取得リクエストの構造体
type NotificationListRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role" validate:"required"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// NotificationResponse 通知レスポンスの構造体
type NotificationResponse struct {
	NotificationID int       `json:"notification_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	TargetRole     string    `json:"target_role"`
	SentAt         time.Time `json:"sent_at"`
	IsRead         bool      `json:"is_read"`
}

// NotificationListResponse 通知一覧レスポンスの構造体
type NotificationListResponse struct {
	Count          int                   `json:"count"`
	Notifications  []NotificationResponse `json:"notifications"`
} 