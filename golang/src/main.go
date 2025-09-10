package main

import (
    "database/sql"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-redis/redis/v8"
)

func main() {
    // MySQL connection
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_NAME"),
    )
    db, err := sql.Open("mysql", dsn)
    
    // Redis connection
    rdb := redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })
}