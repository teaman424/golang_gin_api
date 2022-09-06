package controller

import (
	"gindemo/vender/jwt_auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Auth
// @Router /api/v1/toekn/refresh [post]
// @Accept json
// @Produce json
// @param refreshToken body RefershRequest true "refresh token"
func Refresh(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not access token",
		})
		c.Abort()
		return
	}

	token := strings.Split(auth, "Bearer ")[1]
	result := jwt_auth.Revoke(c, token)
	c.JSON(http.StatusOK, result)
	return
}

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Auth
// @Router /api/v1/toekn/revoke [get]
// @Accept json
// @Produce json
func Revoke(c *gin.Context) {
	// validate request body
	var body struct {
		Token string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newToken, err := jwt_auth.Refresh(c, body.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, newToken)
	return
}
