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
		errorResponse := models.MissingRequiredResponse("X-User-Id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// リクエストボディをパース
	var request models.CreateCourseRequest
	
	// デバッグ用：リクエストボディの内容をログ出力
	bodyBytes, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("DEBUG: Request Body: %s\n", string(bodyBytes))
	
	// ボディを再度設定（io.ReadAllで消費されるため）
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	if err := c.Bind(&request); err != nil {
		errorResponse := models.InvalidFormatResponse("request body", err.Error())
		return c.JSON(http.StatusBadRequest, errorResponse)
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
			errorResponse := models.NewErrorResponse(models.ErrorCodeUnauthorized, models.ErrorMessageUnauthorized, "")
			return c.JSON(http.StatusUnauthorized, errorResponse)
		case errorMsg == "教科情報が存在しません" || strings.HasPrefix(errorMsg, "入力値エラーがあります"):
			errorResponse := models.BadRequestResponse(models.ErrorMessageInvalidInput, errorMsg)
			return c.JSON(http.StatusBadRequest, errorResponse)
		default:
			errorResponse := models.NewErrorResponse(models.ErrorCodeInternalServer, models.ErrorMessageInternalServer, errorMsg)
			return c.JSON(http.StatusInternalServerError, errorResponse)
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
		errorResponse := models.MissingRequiredResponse("course_id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// ユーザーIDをリクエストヘッダーから取得
	userID := c.Request().Header.Get("X-User-Id")
	if userID == "" {
		errorResponse := models.MissingRequiredResponse("X-User-Id")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// リクエストボディをパース
	var request models.UpdateCourseRequest
	
	// デバッグ用：リクエストボディの内容をログ出力
	bodyBytes, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("DEBUG: Request Body: %s\n", string(bodyBytes))
	
	// ボディを再度設定（io.ReadAllで消費されるため）
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	if err := c.Bind(&request); err != nil {
		errorResponse := models.InvalidFormatResponse("request body", err.Error())
		return c.JSON(http.StatusBadRequest, errorResponse)
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
			errorResponse := models.NewErrorResponse(models.ErrorCodeUnauthorized, models.ErrorMessageUnauthorized, "")
			return c.JSON(http.StatusUnauthorized, errorResponse)
		case errorMsg == "教科情報が存在しません" || strings.HasPrefix(errorMsg, "入力値エラーがあります"):
			errorResponse := models.BadRequestResponse(models.ErrorMessageInvalidInput, errorMsg)
			return c.JSON(http.StatusBadRequest, errorResponse)
		case strings.Contains(errorMsg, "you can only update your own courses"):
			errorResponse := models.NewErrorResponse(models.ErrorCodeForbidden, models.ErrorMessageForbidden, "")
			return c.JSON(http.StatusForbidden, errorResponse)
		case strings.Contains(errorMsg, "course not found"):
			errorResponse := models.NewErrorResponse(models.ErrorCodeNotFound, models.ErrorMessageNotFound, "")
			return c.JSON(http.StatusNotFound, errorResponse)
		default:
			errorResponse := models.NewErrorResponse(models.ErrorCodeInternalServer, models.ErrorMessageInternalServer, errorMsg)
			return c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}

	// 授業更新結果をJSON形式で返す
	return c.JSON(http.StatusOK, response)
} 