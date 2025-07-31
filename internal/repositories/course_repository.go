package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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
func (r *CourseRepository) CreateCourse(ctx context.Context, courseData interface{}) error {
	// TODO: 実装
	return nil
} 