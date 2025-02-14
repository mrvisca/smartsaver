package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	RoleId    uint   `gorm:"default:0"`
	Role      Role   `gorm:"foreginKey:RoleId"`
	Name      string `gorm:"type:varchar(200)"`
	Email     string `gorm:"type:varchar(200);unique;not null"`
	Password  string `gorm:"type:varchar(250)"`
	Phone     string `gorm:"type:varchar(25)"`
	Kota      string `gorm:"type:varchar(150)"`
	Provinsi  string `gorm:"type:varchar(150)"`
	Negara    string `gorm:"type:varchar(150)"`
	Statusm   string `gorm:"type:varchar(150)"`
	Pekerjaan string `gorm:"type:varchar(150)"`
	IsActive  bool   `gorm:"type:bool"`
}
