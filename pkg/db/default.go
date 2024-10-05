package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "myAPI/pkg/reader"
    "os"
)

func Default() (*gorm.DB, error) {
    var (
        username string
        password string
        host     string
        port     string
        dbName   string
    )

    reader.GetEnv(".env")
    username = os.Getenv("POSTGRES_USER")
    password = os.Getenv("POSTGRES_PASSWORD")
    host = os.Getenv("POSTGRES_HOST")
    dbName = os.Getenv("POSTGRES_DB")
    port = os.Getenv("POSTGRES_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        host, username, password, dbName, port)

    dbConn, err := gorm.Open(
        postgres.Open(dsn),
        &gorm.Config{
            CreateBatchSize: 500,
        },
    )
    return dbConn, err
}