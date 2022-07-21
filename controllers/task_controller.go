//DBの操作関連の実装
package controllers

import (
	"go-gin-gorm-todo-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ハンドラメソッドのレシーバー用の構造体
type TaskHandler struct {
	//gorm.DBが提供するメソッドを使用可能にするために定義
	Db *gorm.DB
}

//以下、構文
//変数handlerはDB操作の為のポインタレシーバー
//変数cはハンドラの引数wとrを兼ねたようなもの

//レコード（カラム）＝タスク

// タスクの一覧表示
func (handler *TaskHandler) GetAll(c *gin.Context) {
	// レコード(カラム)一覧を格納するため、Task構造体のスライスを変数宣言
	var tasks []models.Task
	// DBから全てのレコードを取得する
	handler.Db.Find(&tasks)
	// index.htmlに全てのレコードを渡す
	c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks})
}

// タスクの新規作成
func (handler *TaskHandler) Create(c *gin.Context) {
	// index.htmlからtextを取得
	text, _ := c.GetPostForm("text")
	// レコードを挿入する
	handler.Db.Create(&models.Task{Text: text})
	//ホーム画面にリダイレクト
	c.Redirect(http.StatusMovedPermanently, "/")
}

// タスクの編集画面
func (handler *TaskHandler) Edit(c *gin.Context) {
	//構造体Taskを簡単に呼び出せるように変数化
	task := models.Task{}
	// index.htmlからidを取得
	id := c.Param("id")
	// idと一致するレコードを取得して第一引数に渡す
	handler.Db.First(&task, id)
	// edit.htmlに編集対象のレコードを渡す
	c.HTML(http.StatusOK, "edit.html", gin.H{"tasks": task})
}

// タスクの更新
func (handler *TaskHandler) Update(c *gin.Context) {
	//構造体Taskを簡単に呼び出せるように変数化
	task := models.Task{}
	// edit.htmlからidを取得
	id := c.Param("id")
	// edit.htmlからtextを取得して左辺に渡す
	text, _ := c.GetPostForm("text")
	// idと一致するレコードを取得して第一引数に渡す
	handler.Db.First(&task, id)
	// textを上書きする
	task.Text = text
	// 指定のレコードを更新する
	handler.Db.Save(&task)

	c.Redirect(http.StatusMovedPermanently, "/")
}

// タスクの削除
func (handler *TaskHandler) Delete(c *gin.Context) {
	//構造体Taskを簡単に呼び出せるように変数化
	task := models.Task{}
	// index.htmlからidを取得
	id := c.Param("id")
	// idと一致するレコードを取得して第一引数に渡す
	handler.Db.First(&task, id)
	// 指定のレコードを削除する
	handler.Db.Delete(&task)
	//ホーム画面にリダイレクト
	c.Redirect(http.StatusMovedPermanently, "/")
}
