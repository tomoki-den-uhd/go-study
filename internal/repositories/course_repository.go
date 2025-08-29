package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tomoki-den-uhd/go-study/internal/models"
)

// CourseRepository 授業リポジトリの構造体
type CourseRepository struct {
	DB *pgxpool.Pool
}

// NewCourseRepository 授業リポジトリのコンストラクタ
func NewCourseRepository(db *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{
		DB: db,
	}
}

// CreateCourse 授業をデータベースに登録する
func (r *CourseRepository) CreateCourse(course *models.Course) (int, error) {
	ctx := context.Background()
	
	// 現在時刻を取得
	now := time.Now()
	
	// SQLクエリを実行
	query := `
		INSERT INTO courses (title, description, teacher_user_id, subject_id, created_at, updated_at, scheduled_at, is_deleted)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING course_id
	`
	
	var courseID int
	err := r.DB.QueryRow(ctx, query,
		course.Title,
		course.Description,
		course.TeacherUserID,
		course.SubjectID,
		now,
		now,
		course.ScheduledAt,
		false, // is_deleted
	).Scan(&courseID)
	
	if err != nil {
		return 0, fmt.Errorf("failed to create course: %w", err)
	}
	
	return courseID, nil
}

// SubjectExists 教科IDの存在確認
func (r *CourseRepository) SubjectExists(subjectID int) (bool, error) {
	ctx := context.Background()
	
	query := `SELECT EXISTS(SELECT 1 FROM subjects WHERE subject_id = $1 AND is_deleted = false)`
	
	var exists bool
	err := r.DB.QueryRow(ctx, query, subjectID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check subject existence: %w", err)
	}
	
	return exists, nil
}

// GetCourseByID 授業IDで授業を取得する
func (r *CourseRepository) GetCourseByID(courseID int) (*models.Course, error) {
	ctx := context.Background()
	
	query := `
		SELECT course_id, title, description, teacher_user_id, subject_id, 
		       created_at, updated_at, scheduled_at, is_deleted
		FROM courses 
		WHERE course_id = $1 AND is_deleted = false
	`
	
	var course models.Course
	err := r.DB.QueryRow(ctx, query, courseID).Scan(
		&course.CourseID,
		&course.Title,
		&course.Description,
		&course.TeacherUserID,
		&course.SubjectID,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.ScheduledAt,
		&course.IsDeleted,
	)
	
	if err != nil {
		return nil, fmt.Errorf("failed to get course: %w", err)
	}
	
	return &course, nil
}

// UpdateCourse 授業を更新する
func (r *CourseRepository) UpdateCourse(courseID int, course *models.Course) error {
	ctx := context.Background()
	
	// 現在時刻を取得
	now := time.Now()
	
	// SQLクエリを実行
	query := `
		UPDATE courses 
		SET title = $1, description = $2, subject_id = $3, 
		    scheduled_at = $4, updated_at = $5
		WHERE course_id = $6 AND is_deleted = false
	`
	
	result, err := r.DB.Exec(ctx, query,
		course.Title,
		course.Description,
		course.SubjectID,
		course.ScheduledAt,
		now,
		courseID,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update course: %w", err)
	}
	
	// 更新された行数を確認
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("course not found or no changes made")
	}
	
	return nil
} 