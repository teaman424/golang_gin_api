package main

import (
	"gindemo/db"
	_ "gindemo/docs"
	md "gindemo/middleware"
	"gindemo/routes"
	"os"
)

/*
*	註解	描述
*	summary	描述該API
*	tags	歸屬同一類的API的tag
*	accept	request的context-type
*	produce	response的context-type
*	param	參數按照 參數名 參數類型 參數的資料類型 是否必須 註解 (中間都要空一格)
*	header	response header return code 參數類型 資料類型 註解
*	router	path httpMethod
 */

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8088
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//db init
	md.Init()
	db.Init()
	router := routes.SetRouter()
	router.Run(":" + os.Getenv("PORT"))
}
