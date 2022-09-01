package main

import (
	_ "gindemo/docs"
	"os"

	"gindemo/controller"
	"gindemo/vender/jwt_auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router := gin.Default()
	router.GET("/balance/", controller.GetBalance)

	demoV1 := router.Group("/demo/v1")

	demoV1.POST("/hello", controller.Hello)

	demoV1.POST("/refresh", controller.Refresh)

	//demoV1.GET("/revoke", controller.Revoke)

	demoV1.Use(jwt_auth.AuthRequired).GET("/hi", controller.Hi)

	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":" + os.Getenv("PORT"))
}
