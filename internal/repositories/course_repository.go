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