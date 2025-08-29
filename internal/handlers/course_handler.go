package handlers

import (
	"bytes"
	"fmt"
	"io"
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
	
	// デバッグ用：リクエストボディの内容をログ出力
	bodyBytes, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("DEBUG: Request Body: %s\n", string(bodyBytes))
	
	// ボディを再度設定（io.ReadAllで消費されるため）
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "入力値エラーがあります: Invalid request body",
			"details": err.Error(),
			"request_body": string(bodyBytes),
			"content_type": c.Request().Header.Get("Content-Type"),
		})
	}
	
	// デバッグ用：パースされた構造体の内容をログ出力
	fmt.Printf("DEBUG: Parsed Request - Title: %s, Description: %s, SubjectID: %d, ScheduledAt: %v\n", 
		request.Title, request.Description, request.SubjectID, request.ScheduledAt)

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

// UpdateCourseHandler 授業更新のハンドラー
func (h *CourseHandler) UpdateCourseHandler(c echo.Context) error {
	// パスパラメータから授業IDを取得
	courseID := c.Param("course_id")
	if courseID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "入力値エラーがあります: course_id parameter is required",
		})
	}

	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "入力値エラーがあります: X-User-Id header is required",
		})
	}

	// リクエストボディをパース
	var request models.UpdateCourseRequest
	
	// デバッグ用：リクエストボディの内容をログ出力
	bodyBytes, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("DEBUG: Request Body: %s\n", string(bodyBytes))
	
	// ボディを再度設定（io.ReadAllで消費されるため）
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "入力値エラーがあります: Invalid request body",
			"details": err.Error(),
			"request_body": string(bodyBytes),
			"content_type": c.Request().Header.Get("Content-Type"),
		})
	}
	
	// デバッグ用：パースされた構造体の内容をログ出力
	fmt.Printf("DEBUG: Parsed Request - Title: %s, Description: %s, SubjectID: %d, ScheduledAt: %v\n", 
		request.Title, request.Description, request.SubjectID, request.ScheduledAt)

	// サービスクラスを呼び出して授業を更新
	response, err := h.courseService.UpdateCourse(courseID, &request, userID)
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
		case strings.Contains(errorMsg, "you can only update your own courses"):
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "アクセス権限がありません",
			})
		case strings.Contains(errorMsg, "course not found"):
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "授業が見つかりません",
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": errorMsg,
			})
		}
	}

	// 授業更新結果をJSON形式で返す
	return c.JSON(http.StatusOK, response)
} 