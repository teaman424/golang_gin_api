package model

type Fruit struct {
	Id          int     `gorm:"column:id" json:"id" example:"1"`
	Category    string  `gorm:"column:category" json:"-" example:"水果"`
	OiginalName string  `gorm:"column:original_name" json:"original_name" example:"蘋果"`
	Name        string  `gorm:"column:name" json:"name" example:"蘋果"`
	Calories    float32 `gorm:"column:calories" json:"calories" example:"20"`
	Quantity    int     `gorm:"column:quantity" json:"-" example:"100"`
}

func (Fruit) TableName() string {
	return "fruit"
}
