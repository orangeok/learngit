package logic

import (
	"fruit_sever/global"
)
import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func SaveUserSql(UserName string, PhoneNumber string, Password string) bool {
	var user User
	err := global.GFruit.Where("phone_number=? ", PhoneNumber).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("gorm err:%s", err)
	}
	if len(user.PhoneNumber) == 0 {
		sql := fmt.Sprintf("INSERT IGNORE INTO `%s`(`user_name`,`phone_number`,`password`,`vip`)VALUES(?,?,?,?)", "user")
		rbq := global.GFruit.Exec(sql, UserName, PhoneNumber, Password, "黄铜")
		if rbq.Error != nil {
			fmt.Println(rbq.Error)
		}
	} else {
		return false
	}
	return true
}
func SaveOrderSql(UserName string, GoodsId string, GoodsNum int64, GoodsName string) {
	sql_c := fmt.Sprintf(`UPDATE %s SET number=number-%d  where id='%s' ;`,
		"commodity",
		GoodsNum,
		GoodsId,
	)
	rbq_c := global.GFruit.Exec(sql_c)
	if rbq_c.Error != nil {
		fmt.Println(rbq_c.Error)
	}
	var orders Orders
	err := global.GFruit.Where("user_name=? AND goods_id=?", UserName, GoodsId).First(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("查询数据库错误:%v", err)
	}
	if len(orders.UserName) == 0 {
		sql := fmt.Sprintf("INSERT IGNORE INTO `%s`(`user_name`,`goods_id`,`goods_num`,`goods_name`)VALUES(?,?,?,?)", "orders")
		rbq := global.GFruit.Exec(sql, UserName, GoodsId, GoodsNum, GoodsName)
		if rbq.Error != nil {
			fmt.Println(rbq.Error)
		}
	} else {
		sql := fmt.Sprintf(`UPDATE %s SET goods_num=goods_num+%d  where user_name='%s' AND goods_id='%s';`,
			"orders",
			GoodsNum,
			UserName,
			GoodsId,
		)
		rdb := global.GFruit.Exec(sql)
		if rdb.Error != nil {
			fmt.Println(rdb.Error)
		}
	}
}
