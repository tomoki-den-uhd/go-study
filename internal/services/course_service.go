package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tomoki-den-uhd/go-study/internal/models"
	"github.com/tomoki-den-uhd/go-study/internal/repositories"
)

// CourseService 授業サービスの構造体
type CourseService struct {
	courseRepo *repositories.CourseRepository
	userService *UserService
}

// NewCourseService 授業サービスのコンストラクタ
func NewCourseService(courseRepo *repositories.CourseRepository, userService *UserService) *CourseService {
	return &CourseService{
		courseRepo: courseRepo,
		userService: userService,
	}
}

// CreateCourse 授業を登録する
func (s *CourseService) CreateCourse(request *models.CreateCourseRequest, userID string) (*models.CreateCourseResponse, error) {
	// ユーザーIDのバリデーション
	userIDInt, err := s.userService.ValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// ユーザーの役割を取得
	userRole, err := s.userService.GetUserRole(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user role: %w", err)
	}

	// 教師のみが授業を作成可能
	if userRole != "teacher" {
		return nil, fmt.Errorf("only teachers can create courses")
	}

	// リクエストのバリデーション（エラーNo. 201）
	if request.Title == "" {
		return nil, fmt.Errorf("入力値エラーがあります: title is required")
	}

	if request.Description == "" {
		return nil, fmt.Errorf("入力値エラーがあります: description is required")
	}

	if request.SubjectID <= 0 {
		return nil, fmt.Errorf("入力値エラーがあります: valid subject_id is required")
	}

	// 教科IDの存在確認（エラーNo. 203）
	subjectExists, err := s.courseRepo.SubjectExists(request.SubjectID)
	if err != nil {
		return nil, fmt.Errorf("データベースエラーが発生しました: %v", err)
	}

	if !subjectExists {
		return nil, fmt.Errorf("教科情報が存在しません")
	}

	// 授業データを作成
	courseData := &models.Course{
		Title:         request.Title,
		Description:   request.Description,
		TeacherUserID: userIDInt,
		SubjectID:     request.SubjectID,
		ScheduledAt:   request.ScheduledAt,
	}

	// リポジトリを呼び出して授業を登録
	courseID, err := s.courseRepo.CreateCourse(courseData)
	if err != nil {
		return nil, fmt.Errorf("データベースエラーが発生しました: %v", err)
	}

	// レスポンスを作成
	response := &models.CreateCourseResponse{
		Status: "OK",
		Info:   map[string]interface{}{},
		Data: models.CourseData{
			CourseID:      courseID,
			TeacherUserID: userIDInt,
			Title:         request.Title,
			Description:   request.Description,
			SubjectID:     request.SubjectID,
			ScheduledAt:   request.ScheduledAt,
			UpdatedAt:     time.Now(),
		},
	}

	return response, nil
}

// UpdateCourse 授業を更新する
func (s *CourseService) UpdateCourse(courseID string, request *models.UpdateCourseRequest, userID string) (*models.UpdateCourseResponse, error) {
	// ユーザーIDのバリデーション
	userIDInt, err := s.userService.ValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// ユーザーの役割を取得
	userRole, err := s.userService.GetUserRole(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user role: %w", err)
	}

	// 教師のみが授業を更新可能
	if userRole != "teacher" {
		return nil, fmt.Errorf("only teachers can update courses")
	}

	// 授業IDの型変換とバリデーション
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, fmt.Errorf("invalid course ID: %w", err)
	}

	if courseIDInt <= 0 {
		return nil, fmt.Errorf("course ID must be positive")
	}

	// 既存の授業を取得
	existingCourse, err := s.courseRepo.GetCourseByID(courseIDInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get course: %w", err)
	}

	// 授業の所有者かチェック
	if existingCourse.TeacherUserID != userIDInt {
		return nil, fmt.Errorf("you can only update your own courses")
	}

	// リクエストのバリデーション（エラーNo. 201）
	if request.Title == "" {
		return nil, fmt.Errorf("入力値エラーがあります: title is required")
	}

	if request.Description == "" {
		return nil, fmt.Errorf("入力値エラーがあります: description is required")
	}

	if request.SubjectID <= 0 {
		return nil, fmt.Errorf("入力値エラーがあります: valid subject_id is required")
	}

	// 教科IDの存在確認（エラーNo. 203）
	subjectExists, err := s.courseRepo.SubjectExists(request.SubjectID)
	if err != nil {
		return nil, fmt.Errorf("データベースエラーが発生しました: %v", err)
	}

	if !subjectExists {
		return nil, fmt.Errorf("教科情報が存在しません")
	}

	// 授業データを更新
	courseData := &models.Course{
		Title:         request.Title,
		Description:   request.Description,
		TeacherUserID: userIDInt,
		SubjectID:     request.SubjectID,
		ScheduledAt:   request.ScheduledAt,
	}

	// リポジトリを呼び出して授業を更新
	err = s.courseRepo.UpdateCourse(courseIDInt, courseData)
	if err != nil {
		return nil, fmt.Errorf("データベースエラーが発生しました: %v", err)
	}

	// 更新後の授業データを取得
	updatedCourse, err := s.courseRepo.GetCourseByID(courseIDInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated course: %w", err)
	}

	// レスポンスを作成
	response := &models.UpdateCourseResponse{
		Status: "OK",
		Info:   map[string]interface{}{},
		Data: models.CourseData{
			CourseID:      updatedCourse.CourseID,
			TeacherUserID: updatedCourse.TeacherUserID,
			Title:         updatedCourse.Title,
			Description:   updatedCourse.Description,
			SubjectID:     updatedCourse.SubjectID,
			ScheduledAt:   updatedCourse.ScheduledAt,
			UpdatedAt:     updatedCourse.UpdatedAt,
		},
	}

	return response, nil
} 