package services

import (
	"fmt"
	"strconv"

	"github.com/tomoki-den-uhd/go-study/internal/models"
	"github.com/tomoki-den-uhd/go-study/internal/repositories"
)

// TestService テストサービスの構造体
type TestService struct {
	testRepo   *repositories.TestRepository
	userService *UserService
}

// NewTestService テストサービスのコンストラクタ
func NewTestService(testRepo *repositories.TestRepository, userService *UserService) *TestService {
	return &TestService{
		testRepo:   testRepo,
		userService: userService,
	}
}

// GetTests 小テストの一覧を取得する
func (s *TestService) GetTests(userID string) ([]models.TestListResponse, error) {
	// ユーザーIDの型変換
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// ユーザーIDから役割を自動取得
	userRole, err := s.userService.GetUserRole(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user role: %w", err)
	}

	// DBアクセス関数を呼ぶ
	tests, err := s.testRepo.SelectTests(userIDInt, userRole)
	if err != nil {
		return nil, fmt.Errorf("failed to get tests: %w", err)
	}

	// データが見つからない場合の処理
	if len(tests) == 0 {
		// 空の配列を返す（エラーではない）
		return []models.TestListResponse{}, nil
	}

	// 必要に応じて結果を整形する
	// 例：日付のフォーマット、スコアの計算など

	// 結果を返す
	return tests, nil
}
