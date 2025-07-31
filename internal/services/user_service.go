package services

import (
	"fmt"
	"strconv"

	"github.com/tomoki-den-uhd/go-study/internal/repositories"
)

// UserService ユーザーサービスの構造体
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService ユーザーサービスのコンストラクタ
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetUserRole ユーザーIDから役割を取得する
func (s *UserService) GetUserRole(userID string) (string, error) {
	// ユーザーIDの型変換
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return "", fmt.Errorf("invalid user ID: %w", err)
	}

	// ユーザーIDのバリデーション
	if userIDInt <= 0 {
		return "", fmt.Errorf("user ID must be positive")
	}

	// データベースから役割を取得
	role, err := s.userRepo.GetUserRole(userIDInt)
	if err != nil {
		return "", fmt.Errorf("failed to get user role: %w", err)
	}

	return role, nil
}

// ValidateUser ユーザーIDの存在確認とバリデーション
func (s *UserService) ValidateUser(userID string) (int, error) {
	// ユーザーIDの型変換
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID: %w", err)
	}

	// ユーザーIDのバリデーション
	if userIDInt <= 0 {
		return 0, fmt.Errorf("user ID must be positive")
	}

	// ユーザーの存在確認
	exists, err := s.userRepo.UserExists(userIDInt)
	if err != nil {
		return 0, fmt.Errorf("failed to validate user: %w", err)
	}

	if !exists {
		return 0, fmt.Errorf("user not found")
	}

	return userIDInt, nil
} 