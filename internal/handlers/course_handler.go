package handlers

import "github.com/labstack/echo/v4"

// CourseHandler 授業ハンドラーの構造体
type CourseHandler struct {
	// TODO: 依存関係を追加
}

// NewCourseHandler 授業ハンドラーのコンストラクタ
func NewCourseHandler() *CourseHandler {
	return &CourseHandler{}
}

// CreateCourseHandler 授業登録のハンドラー
func (h *CourseHandler) CreateCourseHandler(c echo.Context) error {
	// TODO: 実装
	return nil
} 