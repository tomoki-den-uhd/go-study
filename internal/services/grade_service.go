package services

import (
	"fmt"
	"strconv"

	"github.com/tomoki-den-uhd/go-study/internal/models"
	"github.com/tomoki-den-uhd/go-study/internal/repositories"
)

// GradeService 成績サービスの構造体
type GradeService struct {
	gradeRepo   *repositories.GradeRepository
	userService *UserService
}

// NewGradeService 成績サービスのコンストラクタ
func NewGradeService(gradeRepo *repositories.GradeRepository, userService *UserService) *GradeService {
	return &GradeService{
		gradeRepo:   gradeRepo,
		userService: userService,
	}
}

// GetGradeDetail 成績詳細を取得する
func (s *GradeService) GetGradeDetail(gradeID string, userID string) (*models.GradeDetailResponse, error) {
	// ユーザーIDのバリデーション
	userIDInt, err := s.userService.ValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// 成績IDの型変換とバリデーション
	gradeIDInt, err := strconv.Atoi(gradeID)
	if err != nil {
		return nil, fmt.Errorf("invalid grade ID: %w", err)
	}

	if gradeIDInt <= 0 {
		return nil, fmt.Errorf("grade ID must be positive")
	}

	// ユーザーの役割を取得
	userRole, err := s.userService.GetUserRole(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user role: %w", err)
	}

	// 権限チェック（学生は自分の成績のみ、教師は自分のコースの成績のみ）
	if userRole == "student" {
		// 学生の場合：自分の成績のみアクセス可能
		if err := s.validateStudentAccess(gradeIDInt, userIDInt); err != nil {
			return nil, fmt.Errorf("access denied: %w", err)
		}
	} else if userRole == "teacher" {
		// 教師の場合：自分のコースの成績のみアクセス可能
		if err := s.validateTeacherAccess(gradeIDInt, userIDInt); err != nil {
			return nil, fmt.Errorf("access denied: %w", err)
		}
	}
	// その他の役割（admin等）は全ての成績にアクセス可能

	// 成績詳細を取得
	gradeDetail, err := s.gradeRepo.GetGradeDetail(gradeIDInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get grade detail: %w", err)
	}

	return gradeDetail, nil
}

// validateStudentAccess 学生のアクセス権限をチェック
func (s *GradeService) validateStudentAccess(gradeID int, userID int) error {
	// 成績が自分のものかチェック
	grade, err := s.gradeRepo.GetGradeByID(gradeID)
	if err != nil {
		return fmt.Errorf("failed to get grade: %w", err)
	}
	
	if grade.StudentUserID != userID {
		return fmt.Errorf("access denied: you can only access your own grades")
	}
	
	return nil
}

// validateTeacherAccess 教師のアクセス権限をチェック
func (s *GradeService) validateTeacherAccess(gradeID int, userID int) error {
	// 成績が自分のコースのものかチェック
	grade, err := s.gradeRepo.GetGradeByID(gradeID)
	if err != nil {
		return fmt.Errorf("failed to get grade: %w", err)
	}
	
	// 教師がそのコースの担当者かチェック
	isCourseTeacher, err := s.gradeRepo.IsTeacherOfCourse(userID, grade.CourseID)
	if err != nil {
		return fmt.Errorf("failed to check teacher access: %w", err)
	}
	
	if !isCourseTeacher {
		return fmt.Errorf("access denied: you can only access grades from your courses")
	}
	
	return nil
} 