package main

import ()
import (
	"fruit_sever/global"
	"fruit_sever/logic"
)
import (
	"fmt"
	//"os"
)

func main() {
	//	var err_msg string
	init_err := global.InitFruitDb()
	if init_err != nil {
		fmt.Println(init_err)
		//goto EXIT_PROCESS
	}
	logic.RegistRouter()
	/*
		EXIT_PROCESS:
			fmt.Println("初始化异常，退出程序: ", err_msg)
			//fmt.Println(init_err.Error())
			os.Exit(-1)
	*/
}
