//DB用の構造体を定義
package models

import (
	"github.com/jinzhu/gorm"
)

//この部分はgormの定型分
type Task struct {
	gorm.Model
	Text string
}
