package main

import (
	"backend/model"
	"backend/routes"
)

func main() {
	//引用数据库
	model.InitDB()
	model.InitRedis()
	routes.InitRoute()
}
