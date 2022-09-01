package controller

import (
	"gindemo/vender/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Demo
// @Router /demo/v1/refresh [post]
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
