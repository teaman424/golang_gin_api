package controller

import (
	md "gindemo/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Using Refreah Token Update Access Token
// @Success 200 {object} md.AuthToken
// @Tags Auth
// @Router /api/v1/auth/refresh [post]
// @Accept json
// @Produce json
// @param refreshToken body RefershRequest true "refresh token"
func Refresh(c *gin.Context) {
	// validate request body
	var body struct {
		Token string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err.Error(),
		})
		return
	}

	newToken, err := md.Refresh(c, body.Token)
	if err != nil {
		//reply by md.Refresh
		return
	}
	c.JSON(http.StatusOK, newToken)
	return
}

// @Summary Revoke Access Token and Refresh Token
// @Success 200 {object} md.AuthToken
// @Tags Auth
// @Router /api/v1/auth/revoke [get]
// @Accept json
// @Produce json
func Revoke(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": "Not access token",
		})
		c.Abort()
		return
	}

	token := strings.Split(auth, "Bearer ")[1]
	result := md.Revoke(c, token)
	c.JSON(http.StatusOK, result)
	return
}
