package Config

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DbURL() string {
  err := godotenv.Load(".env")

  if err != nil {
    panic("Error loading .env file")
  }

  dbUserName := os.Getenv("DB_USERNAME")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")

  return fmt.Sprintf(
    "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local",
    dbUserName,
    dbPassword,
    dbName,
  )
}
