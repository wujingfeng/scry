package main

import (
	"scry/api"
	"scry/core"
	"scry/global"
)

func main() {
	global.Viper = core.Viper()
	global.Mysql = core.InitMysql()
	api.Question1()
	api.Question2()
	api.Question3()
	api.Question4()
	api.Question5()
	api.Question6()
	api.Question7()
	api.Question2_1()
	api.Question2_2()
	api.Question2_3()
}
