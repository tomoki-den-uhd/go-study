package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tomoki-den-uhd/go-study/internal/models"
)

// GradeRepository 成績リポジトリの構造体
type GradeRepository struct {
	DB *pgxpool.Pool
}

// NewGradeRepository 成績リポジトリのコンストラクタ
func NewGradeRepository(db *pgxpool.Pool) *GradeRepository {
	return &GradeRepository{DB: db}
}

// GetGradeDetail 成績詳細を取得する（2段階取得）
func (g *GradeRepository) GetGradeDetail(gradeID int) (*models.GradeDetailResponse, error) {
	ctx := context.Background()
	
	// 1段階目：成績概要を取得
	gradeSummary, err := g.getGradeSummary(ctx, gradeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get grade summary: %w", err)
	}
	
	// 2段階目：問題詳細を取得
	questionDetails, err := g.getQuestionDetails(ctx, gradeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get question details: %w", err)
	}
	
	// レスポンスを構築
	response := &models.GradeDetailResponse{
		Status: "OK",
		Data: models.GradeData{
			GradeSummary: models.GradeSummary{
				GradeID:         gradeSummary.GradeID,
				StudentUserID:   gradeSummary.StudentUserID,
				TeacherTestID:   gradeSummary.TeacherTestID,
				Score:           gradeSummary.Score,
				Comment:         gradeSummary.Comment,
				TestTitle:       gradeSummary.TestTitle,
				SubjectName:     gradeSummary.SubjectName,
				ScheduledAt:     gradeSummary.ScheduledAt,
				DurationMinutes: gradeSummary.DurationMinutes,
			},
			GradeDetails: questionDetails,
		},
	}
	
	return response, nil
}

// getGradeSummary 成績概要を取得する（設計書1対応）
func (g *GradeRepository) getGradeSummary(ctx context.Context, gradeID int) (*gradeSummary, error) {
	query := `
		SELECT 
			g.grade_id,
			g.student_user_id,
			st.teacher_test_id,
			g.score,
			g.comment,
			tt.title as test_title,
			s.name as subject_name,
			tt.scheduled_at,
			tt.duration_minutes
		FROM grades g
		INNER JOIN student_tests st ON g.student_test_id = st.student_test_id
		INNER JOIN teacher_tests tt ON st.teacher_test_id = tt.teacher_test_id
		INNER JOIN courses c ON g.course_id = c.course_id
		INNER JOIN subjects s ON c.subject_id = s.subject_id
		WHERE g.grade_id = $1 
			AND g.is_deleted = false
			AND st.is_deleted = false
	`
	
	var summary gradeSummary
	err := g.DB.QueryRow(ctx, query, gradeID).Scan(
		&summary.GradeID,
		&summary.StudentUserID,
		&summary.TeacherTestID,
		&summary.Score,
		&summary.Comment,
		&summary.TestTitle,
		&summary.SubjectName,
		&summary.ScheduledAt,
		&summary.DurationMinutes,
	)
	
	if err != nil {
		return nil, fmt.Errorf("failed to scan grade summary: %w", err)
	}
	
	return &summary, nil
}

// GetGradeByID 成績IDで成績を取得する
func (g *GradeRepository) GetGradeByID(gradeID int) (*models.Grade, error) {
	ctx := context.Background()
	
	query := `
		SELECT 
			grade_id,
			student_test_id,
			student_user_id,
			course_id,
			score,
			comment,
			submitted_at,
			is_deleted
		FROM grades
		WHERE grade_id = $1 AND is_deleted = false
	`
	
	var grade models.Grade
	var submittedAt *time.Time
	err := g.DB.QueryRow(ctx, query, gradeID).Scan(
		&grade.ID,
		&grade.StudentTestID,
		&grade.StudentUserID,
		&grade.CourseID,
		&grade.TotalScore,
		&grade.Comment,
		&submittedAt,
		&grade.IsDeleted,
	)
	
	if submittedAt != nil {
		grade.CreatedAt = *submittedAt
		grade.UpdatedAt = *submittedAt
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to get grade by ID: %w", err)
	}
	
	return &grade, nil
}

// IsTeacherOfCourse 教師が指定されたコースの担当者かチェックする
func (g *GradeRepository) IsTeacherOfCourse(teacherUserID int, courseID int) (bool, error) {
	ctx := context.Background()
	
	query := `
		SELECT COUNT(*)
		FROM courses
		WHERE course_id = $1 
			AND teacher_user_id = $2 
			AND is_deleted = false
	`
	
	var count int
	err := g.DB.QueryRow(ctx, query, courseID, teacherUserID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check teacher course access: %w", err)
	}
	
	return count > 0, nil
}

// getQuestionDetails 問題詳細を取得する（設計書2対応）
func (g *GradeRepository) getQuestionDetails(ctx context.Context, gradeID int) ([]models.GradeDetail, error) {
	// まずgradeIDからstudent_test_idを取得
	studentTestID, err := g.getStudentTestIDFromGradeID(ctx, gradeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get student_test_id: %w", err)
	}
	
	query := `
		SELECT 
			tq.test_question_id,
			tq.question_text,
			sta.student_answer,
			sta.is_correct,
			sta.score
		FROM student_test_answers sta
		INNER JOIN test_questions tq ON sta.test_question_id = tq.test_question_id
		INNER JOIN student_tests st ON sta.student_test_id = st.student_test_id
		WHERE sta.student_test_id = $1 
			AND sta.is_deleted = false
		ORDER BY sta.grade_detail_id
	`
	
	rows, err := g.DB.Query(ctx, query, studentTestID)
	if err != nil {
		return nil, fmt.Errorf("failed to query question details: %w", err)
	}
	defer rows.Close()
	
	var details []models.GradeDetail
	for rows.Next() {
		var detail models.GradeDetail
		err := rows.Scan(
			&detail.TestQuestionID,
			&detail.QuestionText,
			&detail.StudentAnswer,
			&detail.IsCorrect,
			&detail.Score,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question detail: %w", err)
		}
		details = append(details, detail)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over question detail rows: %w", err)
	}
	
	return details, nil
}

// getStudentTestIDFromGradeID gradeIDからstudent_test_idを取得する
func (g *GradeRepository) getStudentTestIDFromGradeID(ctx context.Context, gradeID int) (int, error) {
	query := `
		SELECT student_test_id
		FROM grades
		WHERE grade_id = $1 AND is_deleted = false
	`
	
	var studentTestID int
	err := g.DB.QueryRow(ctx, query, gradeID).Scan(&studentTestID)
	if err != nil {
		return 0, fmt.Errorf("failed to get student_test_id from grade_id: %w", err)
	}
	
	return studentTestID, nil
}

// gradeSummary 成績概要の内部構造体
type gradeSummary struct {
	GradeID         int
	StudentUserID   int
	TeacherTestID   int
	Score           int
	Comment         string
	TestTitle       string
	SubjectName     string
	ScheduledAt     time.Time
	DurationMinutes int
} 