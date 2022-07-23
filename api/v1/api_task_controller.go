package v1

import (
	"fmt"
	"go-gin-gorm-todo-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//以下、構文

type ApiTaskHandler struct {
	Db *gorm.DB
}

// 全件取得
// $ curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/ | jsonpp
func (handler *ApiTaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task      // レコード一覧を格納するため、Task構造体のスライスを変数宣言
	handler.Db.Find(&tasks)      // DBから全てのレコードを取得する
	c.JSON(http.StatusOK, tasks) // JSONで全てのレコードを返す
}

// 新規作成
// $ curl -X POST -H "Content-Type: application/json" -d '{"text":"test"}' http://localhost:8080/api/v1/ | jsonpp
func (handler *ApiTaskHandler) Create(c *gin.Context) {
	task := models.Task{}    //構造体を変数化して扱いやすくする
	err := c.BindJSON(&task) //JOSN形式にバインド
	//バインドできなかった時
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//バインドできた時
	handler.Db.Create(&task)     // レコードを挿入する
	c.JSON(http.StatusOK, &task) // JSONで結果を返す
}

// 取得
// $ curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Get(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // リクエストからidを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	c.JSON(http.StatusOK, task) // JSONで取得したレコードを返す
}

// 更新
// $ curl -X PUT -H "Content-Type: application/json" -d '{"text":"update"}' http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Update(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // idを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	request := models.Task{}    // Task構造体の変数宣言
	err := c.BindJSON(&request) //JOSN形式でバインド
	//バインドができなかった時
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//バインドができた時
	task.Text = request.Text     // textを上書きする
	handler.Db.Save(&task)       // 指定のレコードを更新する
	c.JSON(http.StatusOK, &task) // JSONで結果を返す
}

// 削除
// $ curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Delete(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // リクエストからidを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	handler.Db.Delete(&task)    // 指定のレコードを削除する
	msg := fmt.Sprintf("Task [%s] has been deleted.", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
