package service

import (
	"gindemo/db"
	"gindemo/model"
)

func CreateUser(member *model.Member) (err error) {
	mysql := db.GetDb()
	if err = mysql.Table(member.TableName()).Create(member).Error; err != nil {
		return err
	}
	return
}

func CreateUserV2(member *model.Member) (status bool, errMsg string) {
	mysql := db.GetDb()
	result := mysql.Table(member.TableName()).Where("account = ?", member.Account).Find(&model.Member{})
	if result.Error != nil {
		return false, result.Error.Error()
	} else if result.RowsAffected == 1 {
		return false, "account already exist"
	} else {
		if err := mysql.Table(member.TableName()).Create(member).Error; err != nil {
			return false, err.Error()
		}
	}

	return true, ""
}

func CheckUser(account, password string, member *model.Member) (status bool, errMsg string) {
	mysql := db.GetDb()
	result := mysql.Table(member.TableName()).Where("account = ?", account).Where("password = ?", password).First(&member)
	if result.Error != nil {
		return false, result.Error.Error()
	}
	return true, ""
}

func GetUser(userId string, member *model.Member) (err error) {
	mysql := db.GetDb()
	result := mysql.Table(member.TableName()).Where("id = ?", userId).First(&member)
	if result.Error != nil {
		return result.Error
	}
	return
}

func UpdateUserInfo(userId string, info *model.UpdateUser) (err error) {
	mysql := db.GetDb()
	member := model.Member{}
	if err = mysql.Table(member.TableName()).
		Model(member).Where("id = ?", userId).
		Updates(
			map[string]interface{}{
				"name":    info.Name,
				"gender":  info.Gender,
				"phone":   info.Phone,
				"address": info.Address,
			}).Error; err != nil {
		return err
	}
	return
}
