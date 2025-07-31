package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// UserRepository ユーザーリポジトリの構造体
type UserRepository struct {
	DB *pgx.Conn
}

// NewUserRepository ユーザーリポジトリのコンストラクタ
func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserRole ユーザーIDから役割を取得する
func (u *UserRepository) GetUserRole(userID int) (string, error) {
	ctx := context.Background()
	
	query := `
		SELECT role
		FROM users
		WHERE user_id = $1 AND is_deleted = false
	`
	
	var role string
	err := u.DB.QueryRow(ctx, query, userID).Scan(&role)
	if err != nil {
		return "", fmt.Errorf("failed to get user role: %w", err)
	}
	
	return role, nil
}

// UserExists ユーザーIDの存在確認
func (u *UserRepository) UserExists(userID int) (bool, error) {
	ctx := context.Background()
	
	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM users 
			WHERE user_id = $1 AND is_deleted = false
		)
	`
	
	var exists bool
	err := u.DB.QueryRow(ctx, query, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	
	return exists, nil
} 