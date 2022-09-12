package service

import (
	"fmt"
	"gindemo/db"
	"gindemo/model"
	"strconv"
	"strings"
)

func GetUserHabit(userId string, habitInfo *model.Habit) (err error) {
	mysql := db.GetDb()
	tempRuslt := map[string]interface{}{}
	result := mysql.Table(habitInfo.TableName()).Where("member_id = ?", userId).Find(&tempRuslt)
	if result.Error != nil {
		return result.Error
	}

	habitInfo.MemberId = fmt.Sprintf("%v", tempRuslt["member_id"])
	habitInfo.FavoriteList = habitInfo.GetFavoriteStringToSlice(fmt.Sprintf("%v", tempRuslt["favorite_list"]))
	habitInfo.NastyList = habitInfo.GetNastyStringToSlice(fmt.Sprintf("%v", tempRuslt["nasty_list"]))
	return
}

func UpdateUserHabit(userId string, habitInfo *model.UpdateUserHabit) (err error) {
	mysql := db.GetDb()
	habit := model.Habit{}
	result := mysql.Table(habit.TableName()).Where("member_id = ?", userId).Find(map[string]interface{}{})
	fmt.Println("member_id : " + userId)
	fmt.Println("result : " + strconv.FormatInt(result.RowsAffected, 10))
	if result.RowsAffected == 0 {
		if errMsg := mysql.Table(habit.TableName()).Create(map[string]interface{}{
			"member_id":     userId,
			"favorite_list": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(habitInfo.FavoriteList)), ","), "[]"),
			"nasty_list":    strings.Trim(strings.Join(strings.Fields(fmt.Sprint(habitInfo.NastyList)), ","), "[]"),
		}).Error; errMsg != nil {
			return errMsg
		}
	} else {
		if err = mysql.Table(habit.TableName()).
			Model(habit).Where("member_id = ?", userId).
			Updates(
				map[string]interface{}{
					"favorite_list": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(habitInfo.FavoriteList)), ","), "[]"),
					"nasty_list":    strings.Trim(strings.Join(strings.Fields(fmt.Sprint(habitInfo.NastyList)), ","), "[]"),
				}).Error; err != nil {
			return err
		}
	}

	return
}
