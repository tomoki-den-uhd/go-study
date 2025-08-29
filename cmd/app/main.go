package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tomoki-den-uhd/go-study/internal/handlers"
	"github.com/tomoki-den-uhd/go-study/internal/repositories"
	"github.com/tomoki-den-uhd/go-study/internal/services"
)

func main() {
    // .env読み込み
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // 環境変数取得
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // 接続文字列作成
    dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // DB接続（pgxpool使用）
    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer pool.Close()

    fmt.Println("Successfully connected to database!")

    // Echoインスタンスの作成
    e := echo.New()

    // ミドルウェアの追加
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    // 依存関係の注入
    userRepo := repositories.NewUserRepository(pool)
    testRepo := repositories.NewTestRepository(pool)
    gradeRepo := repositories.NewGradeRepository(pool)
    courseRepo := repositories.NewCourseRepository(pool)
    userService := services.NewUserService(userRepo)
    testService := services.NewTestService(testRepo, userService)
    gradeService := services.NewGradeService(gradeRepo, userService)
    courseService := services.NewCourseService(courseRepo, userService)
    testHandler := handlers.NewTestHandler(testService)
    gradeHandler := handlers.NewGradeHandler(gradeService)
    courseHandler := handlers.NewCourseHandler(courseService)

    // ルーティングの設定
    e.GET("/tests", testHandler.GetTestsHandler)
    e.GET("/grades/:grade_id", gradeHandler.GetGradeDetailHandler)
    e.POST("/courses", courseHandler.CreateCourseHandler)
    e.PUT("/courses/:course_id", courseHandler.UpdateCourseHandler)

    // サーバーの起動
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    fmt.Printf("Server starting on port %s...\n", port)
    if err := e.Start(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}