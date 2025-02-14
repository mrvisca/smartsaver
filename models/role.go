package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200)"`
	Desc     string `gorm:"type:text"`
	IsActive bool   `gorm:"type:bool"`
}
