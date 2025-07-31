package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoki-den-uhd/go-study/internal/services"
)

// GradeHandler 成績ハンドラーの構造体
type GradeHandler struct {
	gradeService *services.GradeService
}

// NewGradeHandler 成績ハンドラーのコンストラクタ
func NewGradeHandler(gradeService *services.GradeService) *GradeHandler {
	return &GradeHandler{
		gradeService: gradeService,
	}
}

// GetGradeDetailHandler 成績詳細取得のハンドラー
func (h *GradeHandler) GetGradeDetailHandler(c echo.Context) error {
	// パスパラメータから成績IDを取得
	gradeID := c.Param("grade_id")
	if gradeID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "grade_id parameter is required",
		})
	}

	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "X-User-Id header is required",
		})
	}

	// サービスクラスを呼び出して成績詳細を取得
	gradeDetail, err := h.gradeService.GetGradeDetail(gradeID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// 成績詳細をJSON形式で返す
	return c.JSON(http.StatusOK, gradeDetail)
} 