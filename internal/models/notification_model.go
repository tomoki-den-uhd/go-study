package models

import (
	"time"
)

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