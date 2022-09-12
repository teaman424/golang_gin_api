package controller

import (
	"gindemo/model"
	"gindemo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all friut's information list
// @Success 200 {array} model.Fruit
// @Tags Food
// @Router /api/v1/food/fruit [get]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
func GetFruitList(c *gin.Context) {

	fruitList := &[]model.Fruit{}
	errMsg := service.FruitList(fruitList)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg,
		})
		return
	}

	c.JSON(http.StatusOK, fruitList)
	return
}
