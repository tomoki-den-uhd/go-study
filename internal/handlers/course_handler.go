package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tomoki-den-uhd/go-study/internal/models"
	"github.com/tomoki-den-uhd/go-study/internal/services"
)

// CourseHandler 授業ハンドラーの構造体
type CourseHandler struct {
	courseService *services.CourseService
}

// NewCourseHandler 授業ハンドラーのコンストラクタ
func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

// CreateCourseHandler 授業登録のハンドラー
func (h *CourseHandler) CreateCourseHandler(c echo.Context) error {
	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "入力値エラーがあります: X-User-Id header is required",
		})
	}

	// リクエストボディをパース
	var request models.CreateCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "入力値エラーがあります: Invalid request body",
		})
	}

	// サービスクラスを呼び出して授業を登録
	response, err := h.courseService.CreateCourse(&request, userID)
	if err != nil {
		// エラーメッセージに基づいて適切なHTTPステータスコードを返す
		errorMsg := err.Error()
		
		switch {
		case errorMsg == "認証に失敗しました":
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": errorMsg,
			})
		case errorMsg == "教科情報が存在しません" || strings.HasPrefix(errorMsg, "入力値エラーがあります"):
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errorMsg,
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": errorMsg,
			})
		}
	}

	// 授業登録結果をJSON形式で返す
	return c.JSON(http.StatusCreated, response)
} 