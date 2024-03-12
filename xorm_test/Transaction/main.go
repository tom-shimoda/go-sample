package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Users struct {
	ID   int `xorm:"id"`
	Name string
	Coin int
	Age  int
}

type Items struct {
	ID    int `xorm:"id"`
	Name  string
	Price int
}

type UserHasItems struct {
	ID     int `xorm:"id"`
	UserID int `xorm:"user_id"`
	ItemID int `xorm:"item_id"`
}

func main() {
	//enginを作成します。今回使用するのはDockerで立てたMySQLです。
	engine, err := xorm.NewEngine("mysql", "root:rootpassword@tcp([127.0.0.1]:3306)/dockerdb?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	//buyItemというメソッドを呼び出します
	buyItem(engine)
}

func buyItem(engine *xorm.Engine) {

	// sessionを作成します
	session := engine.NewSession()
	// 最後にsessionを閉じます
	defer session.Close()

	// トランザクションを開始します
	err := session.Begin()

	// ユーザーからコインを減らします
	// コインだけ更新
	user := Users{
		Coin: 80,
	}
	_, err = engine.Where("id=?", 1).Update(&user)
	if err != nil {
		//もし、エラーがあった場合に処理を戻します
		session.Rollback()
		return
	}

	// ユーザーにitemを持たせます
	userItem := UserHasItems{
		ID:     1,
		UserID: 1,
		ItemID: 1,
	}
	_, err = engine.Table("user_has_items").Insert(userItem)
	if err != nil {
		//もし、エラーがあった場合に処理を戻します
		session.Rollback()
		return
	}

	// トランザクションを終了します
	session.Commit()
}
