package routes

import (
	"gindemo/controller"
	md "gindemo/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter() *gin.Engine {

	router := gin.Default()

	//not need middleware api group
	apiV1 := router.Group("/api/v1")

	//need middleware api group
	apiMdV1 := router.Group("/api/v1").Use(md.AuthRequired)

	//Token
	apiV1.POST("/toekn/refresh", controller.Refresh)
	apiV1.GET("/toekn/revoke", controller.Revoke)

	//Users
	apiV1.POST("/users/login", controller.Login)
	apiV1.POST("/users/create", controller.CreateUser)
	apiMdV1.GET("/users/info", controller.GetUser)
	apiMdV1.PATCH("/users/info", controller.UpdateUserInfo)
	apiMdV1.GET("/users/logout", controller.Logout)
	apiMdV1.PATCH("/users/habit", controller.UpdateUserHabit)
	apiMdV1.GET("/users/habit", controller.GetUserHabit)

	//Food
	apiMdV1.GET("/food/fruit", controller.GetFruitList)

	//swagger doc
	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
