package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name     string
	Age      int
	IsActive bool
}

func dbInit() *gorm.DB {
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// 単体作成
func insert(db *gorm.DB) {
	user := User{
		Name:     "太郎",
		Age:      20,
		IsActive: true,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 複数作成
func inserts(db *gorm.DB) {
	users := []User{
		{
			Name:     "花子",
			Age:      25,
			IsActive: true,
		},
		{
			Name:     "龍太郎",
			Age:      30,
			IsActive: false,
		},
		{
			Name:     "太一",
			Age:      35,
			IsActive: false,
		},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 単体取得
func getOne(db *gorm.DB) {
	// 昇順で単体取得
	user1 := User{}
	result1 := db.First(&user1)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println("first:", user1)
	// check error ErrRecordNotFound
	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result1.Error)
	}
	fmt.Println("count:", result1.RowsAffected)

	// 何も指定せず、単体取得
	user2 := User{}
	result2 := db.Take(&user2)
	// SELECT * FROM users LIMIT 1;
	fmt.Println("take:", user2)
	if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result2.Error)
	}

	// 降順で単体取得
	user3 := User{}
	result3 := db.Last(&user3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("last:", user3)
	if errors.Is(result3.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result3.Error)
	}
}

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&User{})

	// --- Create ---
	// insert(db)
	// inserts(db)

	// --- Read ---
    getOne(db)
}
