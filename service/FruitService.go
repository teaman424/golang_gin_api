package service

import (
	"gindemo/db"
	"gindemo/model"
)

func FruitList(fruitList *[]model.Fruit) (err error) {
	mysql := db.GetDb()
	result := mysql.Table(model.Fruit{}.TableName()).Find(fruitList)
	if result.Error != nil {
		return result.Error
	}
	return
}
