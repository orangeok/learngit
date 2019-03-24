package logic

import ()
import ()
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegistRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/fruit_sign_in", GinSignIn)
	router.POST("/fruit_user", GinU)
	router.POST("/fruit_data", GinF)
	router.GET("/fruit_data", GinF)
	router.GET("/fruit_order", GinChangeOrder)
	router.GET("/fruit_get_order", GinGetOrder)
	router.Run(":8000")
}
