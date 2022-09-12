package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	md "gindemo/middleware"
	"gindemo/model"
	"gindemo/service"
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

// @Summary Login User Account
// @Success 200 {object} md.AuthToken
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
			"errMsg": err.Error(),
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
		token, err := md.CreateToken(c, member.Id, member.Account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMsg": err.Error(),
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

// @Summary Create New User
// @Success 200 {object} md.AuthToken
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

	token, err := md.CreateToken(c, nowMember.Id, nowMember.Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, token)
	return
}

// @Summary Get user account information
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

// @Summary Update user account information
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
		c.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
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

// @Summary Logout user account
// @Success 204
// @Tags Users
// @Router /api/v1/users/logout [get]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
func Logout(c *gin.Context) {
	tokenUuid := c.GetString("tokenUuid")
	result := md.Logout(c, tokenUuid)
	c.JSON(http.StatusNoContent, result)
	return
}

// @Summary Get user food habit information
// @Success 200 {object} model.Habit
// @Tags Users
// @Router /api/v1/users/habit [get]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
func GetUserHabit(c *gin.Context) {

	memberHabitInfo := &model.Habit{}
	userId := c.GetString("uuid")

	errMsg := service.GetUserHabit(userId, memberHabitInfo)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberHabitInfo)
	return
}

// @Summary Update user food habit information
// @Success 204
// @Tags Users
// @Router /api/v1/users/habit [patch]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param userHabitInfo body model.UpdateUserHabit true "update user habit info"
func UpdateUserHabit(c *gin.Context) {
	// Validate input
	var input model.UpdateUserHabit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
		return
	}

	userId := c.GetString("uuid")
	fmt.Println("start")
	errMsg := service.UpdateUserHabit(userId, &input)
	fmt.Println("end")
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": errMsg,
		})
		return
	}

	c.JSON(http.StatusNoContent, "true")
	return
}
