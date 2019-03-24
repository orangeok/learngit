package logic

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func GinDatas(c *gin.Context) {
	data := Redis()
	c.JSON(200, gin.H{
		"data": data,
	})
}
func Redis() string {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err == nil {
		fmt.Println("连接成功")
	}
	defer c.Close()
	_, err = c.Do("SET", "username", "you leave is true")
	if err != nil {
		fmt.Println("redis set failed", err)
	}
	username, err := redis.String(c.Do("GET", "username"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return username
}
