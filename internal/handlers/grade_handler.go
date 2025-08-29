package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoki-den-uhd/go-study/internal/models"
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
		errorResponse := models.MissingRequiredResponse("grade_id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		errorResponse := models.MissingRequiredResponse("X-User-Id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// サービスクラスを呼び出して成績詳細を取得
	gradeDetail, err := h.gradeService.GetGradeDetail(gradeID, userID)
	if err != nil {
		errorResponse := models.NewErrorResponse(models.ErrorCodeInternalServer, models.ErrorMessageInternalServer, err.Error())
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	// 成績詳細をJSON形式で返す
	return c.JSON(http.StatusOK, gradeDetail)
} 