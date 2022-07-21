//DBとその接続関連を実装
package db

import (
	"go-gin-gorm-todo-app/models"
	"log"

	"github.com/jinzhu/gorm" //このgormのパッケージでないとgorm.Openが使えない

	_ "github.com/mattn/go-sqlite3"
)

//gormの定型分（DB化した変数でありレシーバーのようなもの）
//グローバルに宣言することで別パッケージからも呼び出せるようにする
var db *gorm.DB

//DBを開く
func Initialize() {
	// 宣言済みのグローバル変数dbをshort variable declaration(:=)で初期化しようとすると
	// ローカル変数dbを初期化することになるので注意する

	// DBのコネクションを接続する
	//gormの定型分
	//第一引数はドライバー、第二引数は作成されるdbファイル
	var err error
	db, err = gorm.Open("sqlite3", "task.db")
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	// ログを有効にする
	db.LogMode(true)

	// Task構造体(Model)を元にマイグレーションを実行する
	db.AutoMigrate(&models.Task{})
}

//DB操作（CRUD）の呼び出し
func Get() *gorm.DB {
	return db
}

// DBのコネクションを切断する
func Close() {
	db.Close()
}
