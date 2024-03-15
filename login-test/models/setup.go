package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() error {
	err := godotenv.Load("env/dev.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}

	driver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)

	DB, err = gorm.Open(driver, dbURI)
	// ↓のように書くとDBはローカル変数として生成されることになりグローバル変数とは無関係となるので注意
    // 参考(https://qiita.com/UHNaKZ/items/637cb3e1c538d8e63ee2)
	// DB, err := gorm.Open(driver, dbURI)

	if err != nil {
		fmt.Printf("Could not connect to the database: %s\n", err.Error())
		return err
	}

	DB.AutoMigrate(&User{})

	return nil
}
