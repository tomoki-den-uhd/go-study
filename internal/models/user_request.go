package models

// UserLoginRequest ユーザーログインリクエストの構造体
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserLoginResponse ユーザーログインレスポンスの構造体
type UserLoginResponse struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token,omitempty"`
	Message  string `json:"message"`
}

// UserRegistrationRequest ユーザー登録リクエストの構造体
type UserRegistrationRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Phone    string `json:"phone"`
	Role     string `json:"role" validate:"required,oneof=student teacher parent"`
}

// UserProfileUpdateRequest ユーザープロフィール更新リクエストの構造体
type UserProfileUpdateRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone"`
} 