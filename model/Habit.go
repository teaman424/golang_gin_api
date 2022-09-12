package model

import (
	"fmt"
	"strings"
)

type Habit struct {
	MemberId     string   `gorm:"column:member_id" json:"member_id" `
	FavoriteList []string `gorm:"column:favorite_list" json:"favorite_list"  example:"櫻桃,釋迦"`
	NastyList    []string `gorm:"column:nasty_list" json:"nasty_list"  example:"櫻桃,釋迦"`
}

type UpdateUserHabit struct {
	FavoriteList []string `gorm:"column:favorite_list" json:"favorite_list" example:"櫻桃,釋迦"`
	NastyList    []string `gorm:"column:nasty_list" json:"nasty_list" example:"櫻桃,釋迦"`
}

func (Habit) TableName() string {
	return "habit"
}

func (Habit) GetFavoriteStringToSlice(s string) []string {
	fmt.Println("s : " + s)
	stringArray := strings.Split(s, ",")
	return stringArray
}

func (Habit) GetNastyStringToSlice(s string) []string {
	stringArray := strings.Split(s, ",")
	return stringArray
}
