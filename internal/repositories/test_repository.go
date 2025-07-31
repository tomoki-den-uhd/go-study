package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tomoki-den-uhd/go-study/internal/models"
)

// TestRepository テストリポジトリの構造体
type TestRepository struct {
	DB *pgxpool.Pool
}

// NewTestRepository テストリポジトリのコンストラクタ
func NewTestRepository(db *pgxpool.Pool) *TestRepository {
	return &TestRepository{DB: db}
}

// SelectTests 小テストの一覧表示をする
// JOINとWHERE条件を使用して、コース、教科、教師の情報も含めて取得
func (t *TestRepository) SelectTests(userID int, userRole string) ([]models.TestListResponse, error) {
	ctx := context.Background()
	
	// デバッグ用ログ
	fmt.Printf("SelectTests called with userID: %d, userRole: %s\n", userID, userRole)
	
	// ユーザーの役割に応じてクエリを変更
	var query string
	var args []interface{}
	
	if userRole == "teacher" {
		// 教師の場合：自分が作成したテストを取得
		query = `
			SELECT 
				tt.teacher_test_id,
				tt.title,
				tt.description,
				tt.duration_minutes,
				c.title as course_title,
				s.name as subject_name,
				u.name as teacher_name,
				tt.scheduled_at,
				tt.is_draft,
				tt.created_at
			FROM teacher_tests tt
			JOIN courses c ON tt.course_id = c.course_id
			JOIN subjects s ON c.subject_id = s.subject_id
			JOIN users u ON c.teacher_user_id = u.user_id
			WHERE tt.created_by = $1 
				AND tt.is_deleted = false
				AND c.is_deleted = false
				AND s.is_deleted = false
				AND u.is_deleted = false
			ORDER BY tt.created_at DESC
		`
		args = []interface{}{userID}
	} else {
		// 学生やその他の役割の場合：学生が受講しているコースのテストのみ取得
		query = `
			SELECT DISTINCT
				tt.teacher_test_id,
				tt.title,
				tt.description,
				tt.duration_minutes,
				c.title as course_title,
				s.name as subject_name,
				u.name as teacher_name,
				tt.scheduled_at,
				tt.is_draft,
				tt.created_at
			FROM teacher_tests tt
			JOIN courses c ON tt.course_id = c.course_id
			JOIN subjects s ON c.subject_id = s.subject_id
			JOIN users u ON c.teacher_user_id = u.user_id
			JOIN attendances a ON c.course_id = a.course_id
			WHERE a.student_user_id = $1 
				AND tt.is_deleted = false
				AND c.is_deleted = false
				AND s.is_deleted = false
				AND u.is_deleted = false
				AND a.is_deleted = false
				AND tt.is_draft = false
			ORDER BY tt.scheduled_at ASC
		`
		args = []interface{}{userID}
		
		// デバッグ用：クエリ実行前のログ
		fmt.Printf("Executing query for student/other role (userID: %d)\n", userID)
		fmt.Printf("Query: %s\n", query)
		fmt.Printf("Args: %v\n", args)
	}

	rows, err := t.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tests: %w", err)
	}
	defer rows.Close()

	var tests []models.TestListResponse
	for rows.Next() {
		var test models.TestListResponse
		err := rows.Scan(
			&test.TeacherTestID,
			&test.Title,
			&test.Description,
			&test.DurationMinutes,
			&test.CourseTitle,
			&test.SubjectName,
			&test.TeacherName,
			&test.ScheduledAt,
			&test.IsDraft,
			&test.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan test row: %w", err)
		}
		tests = append(tests, test)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over test rows: %w", err)
	}

	return tests, nil
}
