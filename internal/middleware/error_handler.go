package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tomoki-den-uhd/go-study/internal/models"
)

// ErrorHandler カスタムエラーハンドラー
func ErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else {
		msg = err.Error()
	}

	// エラーログの出力
	fmt.Printf("[ERROR] %s %s - %d: %v\n", 
		c.Request().Method, 
		c.Request().URL.Path, 
		code, 
		msg,
	)

	// 共通エラーレスポンス形式で返す
	errorResponse := models.NewErrorResponse(code, fmt.Sprintf("%v", msg), "")
	c.JSON(code, errorResponse)
}

// RequestLogger リクエストログミドルウェア
func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			
			// リクエスト開始ログ
			fmt.Printf("[INFO] %s %s - Started\n", 
				c.Request().Method, 
				c.Request().URL.Path,
			)
			
			err := next(c)
			
			// リクエスト完了ログ
			duration := time.Since(start)
			status := c.Response().Status
			
			fmt.Printf("[INFO] %s %s - %d - %v\n", 
				c.Request().Method, 
				c.Request().URL.Path, 
				status, 
				duration,
			)
			
			return err
		}
	}
}

// ValidationErrorHandler バリデーションエラーハンドラー
func ValidationErrorHandler(err error) models.CommonResponse {
	return models.NewErrorResponse(
		models.ErrorCodeInvalidInput,
		models.ErrorMessageInvalidInput,
		err.Error(),
	)
} 