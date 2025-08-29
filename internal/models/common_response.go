package models

import "fmt"

// CommonResponse 共通レスポンス形式
type CommonResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data,omitempty"`
	Error   *ErrorResponse         `json:"error,omitempty"`
	Info    map[string]interface{} `json:"info,omitempty"`
}

// ErrorResponse エラーレスポンス形式
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// SuccessResponse 成功レスポンスのヘルパー関数
func SuccessResponse(data interface{}, message string) CommonResponse {
	return CommonResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse エラーレスポンスのヘルパー関数
func NewErrorResponse(code int, message string, details string) CommonResponse {
	return CommonResponse{
		Status: "error",
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
}

// BadRequestResponse 400 Bad Requestエラーレスポンスのヘルパー関数
func BadRequestResponse(message string, details string) CommonResponse {
	return NewErrorResponse(ErrorCodeInvalidInput, message, details)
}

// MissingRequiredResponse 必須項目不足エラーレスポンスのヘルパー関数
func MissingRequiredResponse(field string) CommonResponse {
	return NewErrorResponse(
		ErrorCodeMissingRequired,
		ErrorMessageMissingRequired,
		fmt.Sprintf("必須項目 '%s' が不足しています", field),
	)
}

// InvalidFormatResponse 形式エラーレスポンスのヘルパー関数
func InvalidFormatResponse(field string, details string) CommonResponse {
	return NewErrorResponse(
		ErrorCodeInvalidFormat,
		ErrorMessageInvalidFormat,
		fmt.Sprintf("項目 '%s' の形式が正しくありません: %s", field, details),
	)
}

// エラーコード定数
const (
	// 入力値エラー (400系)
	ErrorCodeInvalidInput     = 400
	ErrorCodeMissingRequired  = 400
	ErrorCodeInvalidFormat    = 400
	
	// 認証・認可エラー (401-403系)
	ErrorCodeUnauthorized     = 401
	ErrorCodeForbidden        = 403
	
	// リソースエラー (404系)
	ErrorCodeNotFound         = 404
	ErrorCodeResourceNotFound = 404
	
	// サーバーエラー (500系)
	ErrorCodeInternalServer   = 500
	ErrorCodeDatabaseError    = 500
)

// エラーメッセージ定数
const (
	ErrorMessageInvalidInput     = "入力値エラーがあります"
	ErrorMessageMissingRequired  = "必須項目が不足しています"
	ErrorMessageInvalidFormat    = "形式が正しくありません"
	ErrorMessageUnauthorized     = "認証に失敗しました"
	ErrorMessageForbidden        = "アクセス権限がありません"
	ErrorMessageNotFound         = "リソースが見つかりません"
	ErrorMessageInternalServer   = "サーバー内部エラーが発生しました"
	ErrorMessageDatabaseError    = "データベースエラーが発生しました"
) 