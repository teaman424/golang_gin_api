package routes

import (
	"gindemo/controller"
	"gindemo/vender/jwt_auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter() *gin.Engine {

	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	apiV1.GET("/demo/balance/", controller.GetBalance)

	apiV1.POST("/users/login", controller.Login)
	apiV1.POST("/users/create", controller.CreateUser)

	apiV1.POST("/toekn/refresh", controller.Refresh)

	apiV1.GET("/toekn/revoke", controller.Revoke)

	apiV1.Use(jwt_auth.AuthRequired).GET("/demo/name", controller.GetName)
	apiV1.Use(jwt_auth.AuthRequired).GET("/users/info", controller.GetUser)
	apiV1.Use(jwt_auth.AuthRequired).PATCH("/users/info", controller.UpdateUserInfo)

	//swagger doc
	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
