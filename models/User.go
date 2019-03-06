package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Address  string `gorm:"type:varchar(255)"`
	Phone    string `gorm:"type:varchar(11)"`
	UserName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Hobbies  string `gorm:"type:varchar(255)"`
	About    string `form:"type:text"`
	PassWord []byte
	Token    []byte
	Block    int8
}
