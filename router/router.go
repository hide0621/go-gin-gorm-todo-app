//ルーティングの実装
package router

import (
	"go-gin-gorm-todo-app/controllers"
	"go-gin-gorm-todo-app/db"

	"github.com/gin-gonic/gin"
)

func Router() {
	//以下、構文

	// gin内で定義されているEngine構造体インスタンスを取得
	// Router、HTML Rendererの機能を内包している
	router := gin.Default()

	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	// 今回のケースではtemplatesディレクトリ配下のhtmlファイルを関連付けている
	router.LoadHTMLGlob("templates/*.html")

	// TaskHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.TaskHandler{
		Db: db.Get(),
	}

	// 各パスにGET/POSTメソッドでリクエストされた時のレスポンスを登録する
	// 第一引数にパス、第二引数にHandlerを登録する
	router.GET("/", handler.GetAll)            // 一覧表示
	router.POST("/", handler.Create)           // 新規作成
	router.GET("/:id", handler.Edit)           // 編集画面
	router.POST("/update/:id", handler.Update) // 更新
	router.POST("/delete/:id", handler.Delete) // 削除

	// Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する（サーバーが起動される）
	router.Run()
}
