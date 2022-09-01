package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 1000

// @Success 200 {string} string
// @Tags Demo
// @Router /demo/v1/hi [get]
// @Security ApiKeyAuth
func GetName(c *gin.Context) {
	account, ok := c.Get("account")
	if !ok {
		account = "No Name"
	}
	c.JSON(200, "Hi "+account.(string))
}

// @Success 200 {string} string
// @Router /balance/ [get]
func GetBalance(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})
}
