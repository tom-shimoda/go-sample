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

// ----------------------------------------
// Insert
// ----------------------------------------
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

// ----------------------------------------
// Read
// ----------------------------------------
// 単体取得
func getOne(db *gorm.DB) {
	// プライマリーキーの昇順で単体取得 (プライマリーキーがない場合は最初のフィールドでソートされる)
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

	// プライマリーキーの降順で単体取得 (プライマリーキーがない場合は最初のフィールドでソートされる)
	user3 := User{}
	result3 := db.Last(&user3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("last:", user3)
	if errors.Is(result3.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result3.Error)
	}
}

// 全件取得
func find(db *gorm.DB) {
	users := []User{}
	result := db.Find(&users)
	fmt.Println("user:", users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// ----------------------------------------
// Update
// ----------------------------------------
// 更新(upsert)
func save(db *gorm.DB) {
	// 構造体にidが無い場合はinsertされる
	user1 := User{}
	user1.Name = "花子"
	result1 := db.Save(&user1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println("count:", result1.RowsAffected)
	fmt.Println("user1:", user1)

	// 先にユーザーを取得する
	user2 := User{}
	db.First(&user2)

	// 構造体にidがある場合はupdateされる
	user2.Name = "たけし"
	result2 := db.Save(&user2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println("count:", result2.RowsAffected)
	fmt.Println("user2:", user2)
}

// 単一のカラムを更新する
// (db.Model().~で操作しているが、Model側に主キーが設定されているためである。
//
//	db.Where("id=?")とするとフィルタ結果が該当なしとなってしまう)
func update(db *gorm.DB) {
	result := db.Model(&User{}).Where("id = 3").Update("name", "ジョージ")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	user := User{}
	db.Where("id = 3").Take(&user)
	fmt.Println("user:", user)
}

// 複数のカラムを更新する
// (db.Model().~で操作しているが、Model側に主キーが設定されているためである。
//
//	db.Where("id=?")とするとフィルタ結果が該当なしとなってしまう)
func updates(db *gorm.DB) {
	result := db.Model(&User{}).Where("id = 1").Updates(User{Name: "Taro", Age: 10, IsActive: true})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	user := User{}
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)
}

// ゼロ値更新されない
func noUpdates(db *gorm.DB) {
	// IsActiveカラムは更新されない
	result := db.Model(User{}).Where("id = 1").Updates(User{Name: "マリオ", IsActive: false})
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("No update:::")
	user := User{}
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)

	// Selectで指定することで更新されます
	result = db.Model(User{}).Where("id = 1").Select("name", "is_active").Updates(User{Name: "マリオ", IsActive: false})

	fmt.Println("Select update:::")
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)
}

// ----------------------------------------
// Delete
// ----------------------------------------
func delete(db *gorm.DB) {
    // 論理削除
	db.Where("id = 1").Delete(&User{})

    // 物理削除
    // db.Unscoped().Where("id = 1").Delete(&User{})
}

// ----------------------------------------
// main
// ----------------------------------------
func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&User{})

	// --- Create ---
	// insert(db) // 一件追加
	// inserts(db) // 複数件追加

	// --- Read ---
	// getOne(db) // 一件読み取り
	// find(db) // 全件読み取り

	// --- Update ---
	// save(db) // 既に存在すれば更新。なければ追加
	// update(db) // 単一カラム更新
	// updates(db) // 複数カラム更新
	// noUpdates(db) // 0値設定の場合更新されないので注意

	// --- Delete ---
    // delete(db) // 一件削除
}
