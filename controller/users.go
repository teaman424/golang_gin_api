package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gindemo/model"
	"gindemo/service"
	"gindemo/vender/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Account  string `json:"account" example:"Jack"`
	Password string `json:"password" example:"12345"`
}

type RefershRequest struct {
	Token string `json:"token" example:"xxcjdjfcidasjcodioi"`
}

type UpdateUserRequest struct {
	model.UpdateUser
}

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Users
// @Router /api/v1/users/login [post]
// @Accept json
// @Produce json
// @param userInfo body LoginRequest true "Add account"
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
	//using md5 encoder password
	signByte := []byte(body.Password)
	hash := md5.New()
	hash.Write(signByte)
	md5Password := hex.EncodeToString(hash.Sum(nil))

	// check account and password is correct
	member := &model.Member{}
	result, errMsg := service.CheckUser(body.Account, md5Password, member)
	if result {
		fmt.Println("member.Id : " + member.Id)
		fmt.Println("member.Account : " + member.Account)
		token, err := jwt_auth.CreateToken(c, member.Id, member.Account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, token)
		return
	} else {
		// incorrect account or password
		c.JSON(http.StatusUnauthorized, gin.H{
			"errMsg": errMsg,
		})
		return
	}

}

// @Success 200 {object} jwt_auth.AuthToken
// @Tags Users
// @Router /api/v1/users/create [post]
// @Accept json
// @Produce json
// @param account body LoginRequest true "Add account"
func CreateUser(c *gin.Context) {
	// validate request body
	var input LoginRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err.Error(),
		})
		return
	}
	//  using md5 encoder password
	signByte := []byte(input.Password)
	hash := md5.New()
	hash.Write(signByte)
	md5Password := hex.EncodeToString(hash.Sum(nil))
	// fmt.Println("DB Init Get DB start")
	// memberTableCont := db.GetDb()
	// fmt.Println("DB Init Get DB end")

	userUuid := hex.EncodeToString(uuid.New().NodeID())

	nowMember := model.Member{
		Id:              userUuid,
		Account:         input.Account,
		Password:        md5Password,
		IsVerify:        false,
		VerifyCode:      "",
		VerifyExpiresAt: "0",
		Name:            "",
		Gender:          "nil",
		Phone:           "",
		Address:         "",
	}

	fmt.Println("DB start Using")
	fmt.Println(nowMember)
	// if err = service.CreateUser(&nowMember); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	status, errMsg := service.CreateUserV2(&nowMember)
	if !status {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg,
		})
		return
	}

	token, err := jwt_auth.CreateToken(c, nowMember.Id, nowMember.Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, token)
	return
}

// @Success 200 {object} model.Member
// @Tags Users
// @Router /api/v1/users/info [get]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {

	memberInfo := &model.Member{}
	userId := c.GetString("uuid")

	errMsg := service.GetUser(userId, memberInfo)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg,
		})
		return
	}

	c.JSON(http.StatusOK, memberInfo)
	return
}

// @Success 204
// @Tags Users
// @Router /api/v1/users/info [patch]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param userInfo body model.UpdateUser true "update user info"
func UpdateUserInfo(c *gin.Context) {
	// Validate input
	var input model.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("uuid")
	errMsg := service.UpdateUserInfo(userId, &input)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg,
		})
		return
	}

	c.JSON(http.StatusNoContent, "true")
	return
}
