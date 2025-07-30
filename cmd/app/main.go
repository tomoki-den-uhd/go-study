package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/jackc/pgx/v5"
)

func main() {
    // .env読み込み
    err := godotenv.Load("../../.env")
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

    // DB接続
    conn, err := pgx.Connect(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer conn.Close(context.Background())

    fmt.Println("Successfully connected to database!")

    // ここでクエリ等を実行可能
}