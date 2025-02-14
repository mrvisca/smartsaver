package settings

import (
	"smartsaver/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@(localhost)/smartsaver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{}).AddForeignKey("role_id", "roles(id)", "NO ACTION", "NO ACTION")
}
