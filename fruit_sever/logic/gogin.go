package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
import (
	"fruit_sever/global"
)
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func GinF(c *gin.Context) {
	id := c.Query("id")
	data := SelectSql(id)
	c.JSON(200, gin.H{
		"data": data,
	})
}
func GinSignIn(c *gin.Context) {
	var data bool
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	var si Sign_In
	fmt.Println(string(body))
	err = json.Unmarshal(body, &si)
	fmt.Println(err, si)
	user_name := si.NickName       //c.Query("user_name")
	phone_number := si.PhoneNumber //c.Query("phone_number")
	password := si.Password        //c.Query("password")
	//phone_numberx, err := strconv.ParseInt(phone_number, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PASS", password)
	if password != "" || user_name != "" {
		data = SaveUserSql(user_name, phone_number, password)
	}
	if data {
		c.JSON(200, gin.H{
			"data": "1",
		})
	} else {
		c.JSON(200, gin.H{
			"data": "2",
		})
	}
}
func GinU(c *gin.Context) {
	var data User
	//	phone_number := c.Query("phone_number")
	//	password := c.Query("password")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	var uk UserKey
	json.Unmarshal(body, &uk)
	fmt.Println(string(body))
	fmt.Println(err, uk)
	phone_number := uk.PhoneNumber
	password := uk.Password
	fmt.Println("PASS", password)
	if phone_number != "" || password != "" {
		data = UserSql(phone_number, password)
	}
	if data.UserName == "" {
		c.JSON(200, gin.H{
			"data": 2,
		})
	} else {
		c.JSON(200, gin.H{
			"data": data.UserName,
		})
	}
}
func GinChangeOrder(c *gin.Context) {
	user_name := c.Query("user_name")
	goods_id := c.Query("goods_id")
	goods_num := c.Query("goods_num")
	goods_name := c.Query("goods_name")
	goods_numx, err := strconv.ParseInt(goods_num, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PASS", goods_id)
	SaveOrderSql(user_name, goods_id, goods_numx, goods_name)
	if goods_id != "" && goods_numx != 0 && user_name != "" && goods_name != "" {

	}
	if goods_id == "" {
		c.JSON(200, gin.H{
			"data": "密码错误或用户不存在",
		})
	} else {
		c.JSON(200, gin.H{
			"data": goods_id,
		})
	}
}
func GinGetOrder(c *gin.Context) {
	var data []Orders
	user_name := c.Query("user_name")
	if user_name != "" {
		data = OrderSql(user_name)
	}
	c.JSON(200, gin.H{
		"data": data,
	})
}
func GoGin() {

}
func SelectSql(n string) Commodity {
	id, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	var c Commodity
	err3 := global.GFruit.Where("id=?", id).First(&c).Error
	if err3 != nil && err3 != gorm.ErrRecordNotFound {
		fmt.Printf("gorm err:%v", err3)
	}
	return c
}
func UserSql(phone_number string, password string) User {

	var u User
	err := global.GFruit.Where("phone_number=? AND password=?", phone_number, password).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("gorm err:%v", err)
	}
	fmt.Println(u, password)
	return u
}
func OrderSql(user_name string) []Orders {
	var o []Orders
	err := global.GFruit.Where("user_name=? ", user_name).Find(&o).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("gorm err:%v", err)
	}
	return o
}
