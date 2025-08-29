package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoki-den-uhd/go-study/internal/models"
	"github.com/tomoki-den-uhd/go-study/internal/services"
)

// TestHandler テストハンドラーの構造体
type TestHandler struct {
	testService *services.TestService
}

// NewTestHandler テストハンドラーのコンストラクタ
func NewTestHandler(testService *services.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

// GetTestsHandler 小テスト一覧取得のハンドラー
func (h *TestHandler) GetTestsHandler(c echo.Context) error {
	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		errorResponse := models.MissingRequiredResponse("X-User-Id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// サービスクラスを呼び出してテスト一覧を取得
	tests, err := h.testService.GetTests(userID)
	if err != nil {
		errorResponse := models.NewErrorResponse(models.ErrorCodeInternalServer, models.ErrorMessageInternalServer, err.Error())
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	// テストの一覧をJSON形式で返す
	return c.JSON(http.StatusOK, map[string]interface{}{
		"tests": tests,
		"count": len(tests),
	})
}
