package controller

import (
	"gindemo/vender/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloRequest struct {
	Account  string `json:"account" example:"Jack"`
	Password string `json:"password" example:"12345"`
}

type RefershRequest struct {
	Token string `json:"token" example:"xxcjdjfcidasjcodioi"`
}

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Demo
// @Router /demo/v1/hello [post]
// @Accept json
// @Produce json
// @param account body HelloRequest true "Add account"
func Login(c *gin.Context) {
	// validate request body
	var body struct {
		Account  string
		Password string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// check account and password is correct
	if body.Account == "Kenny" && body.Password == "123456" {
		token, err := jwt_auth.CreateToken(c, body.Account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, token)
		return
	}
	// incorrect account or password
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	return
}
