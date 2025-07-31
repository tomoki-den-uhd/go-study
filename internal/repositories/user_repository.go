package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository ユーザーリポジトリの構造体
type UserRepository struct {
	DB *pgxpool.Pool
}

// NewUserRepository UserRepositoryのコンストラクタ
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// ExistsUser ユーザーが存在するかチェック
func (r *UserRepository) ExistsUser(userID int) (bool, error) {
	ctx := context.Background()
	
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users 
			WHERE user_id = $1 AND is_deleted = false
		)
	`
	
	var exists bool
	err := r.DB.QueryRow(ctx, query, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	
	return exists, nil
}

// GetUserRole ユーザーの役割を取得
func (r *UserRepository) GetUserRole(userID int) (string, error) {
	ctx := context.Background()
	
	query := `
		SELECT role
		FROM users
		WHERE user_id = $1 AND is_deleted = false
	`
	
	var role string
	err := r.DB.QueryRow(ctx, query, userID).Scan(&role)
	if err != nil {
		return "", fmt.Errorf("failed to get user role: %w", err)
	}
	
	return role, nil
} 