package controller

import (
	"gindemo/vender/jwt_auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 1000

// type Response struct {
// 	AccessToken string `json:"token" example:"sdjjdshjduiashuyehcuisco=="`
// 	TokenType   string `json:"tokenType" example:"Bearer"`
// }

var jwtSecret = []byte("secret")

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
func Hello(c *gin.Context) {
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

// @Success 200 {string} string
// @Tags Demo
// @Router /demo/v1/hi [get]
// @Security ApiKeyAuth
func Hi(c *gin.Context) {
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
