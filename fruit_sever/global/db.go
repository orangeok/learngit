package global

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
import ()
import (
	"fmt"
)

var GFruit *gorm.DB = nil

func InitFruitDb() error {
	var err error
	for ok := true; ok; ok = false {
		GFruit, err = gorm.Open("mysql", "root:jffldzkd@(134.175.42.55)/fruit?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
		if err != nil {
			break
		}
		GFruit.DB().SetMaxOpenConns(20)
		GFruit.DB().SetMaxIdleConns(20)
		GFruit.LogMode(true)
		//GFruit.SetLogger("")
		fmt.Println("初始化PigDb成功")

		return nil
	}
	return err
}
func CloseGFruitDb() {

	if GFruit != nil {
		GFruit.Close()
	}
}
