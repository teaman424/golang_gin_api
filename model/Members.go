package model

type Member struct {
	Id              string `gorm:"column:id" json:"id" `
	Account         string `gorm:"column:account" json:"account"`
	Password        string `gorm:"column:password" json:"-"`
	IsVerify        bool   `gorm:"column:isVerify" json:"isVerify"`
	VerifyCode      string `gorm:"column:verifyCode" json:"-"`
	VerifyExpiresAt string `gorm:"column:verifyExpiresAt" json:"-"`
	Name            string `gorm:"column:name" json:"name"`
	Gender          string `gorm:"column:gender" json:"gender"`
	Phone           string `gorm:"column:phone" json:"phone"`
	Address         string `gorm:"column:address" json:"address"`
}

type UpdateUser struct {
	Name    string `json:"name" example:"Jack"`
	Address string `json:"address" example:"地球"`
	Phone   string `json:"phone" example:"0987654321"`
	Gender  string `json:"gender" example:"m"`
}

func (Member) TableName() string {
	return "members"
}
