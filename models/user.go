package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	RoleId     uint   `gorm:"default:0"`
	Role       Role   `gorm:"foreginKey:RoleId"`
	Name       string `gorm:"type:varchar(200)"`
	Email      string `gorm:"type:varchar(200);unique;not null"`
	Password   string `gorm:"type:varchar(250)"`
	Jenkel     string `gorm:"type:varchar(100);not null"`
	Phone      string `gorm:"type:varchar(25)"`
	Nik        string `gorm:"type:varchar(50);unique"`
	Kota       string `gorm:"type:varchar(150)"`
	Provinsi   string `gorm:"type:varchar(150)"`
	Negara     string `gorm:"type:varchar(150)"`
	Statusm    string `gorm:"type:varchar(150)"`
	Pekerjaan  string `gorm:"type:varchar(150)"`
	Kode       string `gorm:"type:varchar(50)"`
	JoinDate   string `gorm:"type:varchar(150)"`
	IsVerified bool   `gorm:"type:bool"`
	IsActive   bool   `gorm:"type:bool"`
}

type ResRegister struct {
	ID         uint   `json:"id"`
	RoleId     uint   `json:"role_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Jenkel     string `json:"jenkel"`
	Phone      string `json:"phone"`
	Nik        string `json:"nik"`
	Kota       string `json:"kota"`
	Provinsi   string `json:"provinsi"`
	Negara     string `json:"negara"`
	Statusm    string `json:"statusm"`
	Pekerjaan  string `json:"pekerjaan"`
	Kode       string `json:"kode"`
	JoinDate   string `json:"join_date"`
	IsVerified bool   `json:"is_verified"`
	IsActive   bool   `json:"is_active"`
}

type ProfilePengguna struct {
	ID         uint   `json:"id"`
	RoleId     uint   `json:"role_id"`
	RoleName   string `json:"role_name"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Jenkel     string `json:"jenkel"`
	Phone      string `json:"phone"`
	Nik        string `json:"nik"`
	Kota       string `json:"kota"`
	Provinsi   string `json:"provinsi"`
	Negara     string `json:"negara"`
	Statusm    string `json:"statusm"`
	Pekerjaan  string `json:"pekerjaan"`
	Kode       string `json:"kode"`
	JoinDate   string `json:"join_date"`
	IsVerified bool   `json:"is_verified"`
	IsActive   bool   `json:"is_active"`
}
