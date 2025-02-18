package seeders

import (
	"log"
	"smartsaver/helpers"
	"smartsaver/models"
	"smartsaver/settings"

	"github.com/gin-gonic/gin"
)

func RoleSeed(c *gin.Context) {
	// Variabel model role
	modRole := []models.Role{}

	// Variabel hitung record
	var hitRole int64
	settings.DB.Model(&modRole).Count(&hitRole)

	// Kondisi jika role belum terisi
	if hitRole > 0 {
		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Role Aplikasi Siap Digunakan!")
	} else {
		// Struct simpan data role seeder
		admRole := models.Role{
			Name:     "Administrator",
			Desc:     "Administrator Aplikasi",
			IsActive: true,
		}

		penggunaRole := models.Role{
			Name:     "Pengguna",
			Desc:     "Pengguna Aplikasi",
			IsActive: true,
		}

		// Simpan struct data
		settings.DB.Create(&admRole)
		settings.DB.Create(&penggunaRole)

		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Role Aplikasi Sukses Dibuat!")
	}
}

func UserSeed(c *gin.Context) {
	// Variabel model user
	modUser := []models.User{}

	// Variabel hitung data pengguna
	var hitUser int64
	settings.DB.Model(&modUser).Count(&hitUser)

	// Kondisi jika data user sudah terisi
	if hitUser > 0 {
		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Pengguna Aplikasi Siap Digunakan!")
	} else {
		// Variabel password
		passAdm := "11223344"

		// Enkripsi password
		pass1, _ := helpers.HashPassword(passAdm)

		// Struct buat data admin
		simAdm := models.User{
			RoleId:     1,
			Name:       "Visca Putra",
			Email:      "mrvisca2018@gmail.com",
			Password:   pass1,
			Jenkel:     "Laki-Laki",
			Phone:      "6282140466335",
			Nik:        "3518110710950002",
			Kota:       "Nganjuk",
			Provinsi:   "Jawa Timur",
			Negara:     "Indonesia",
			Statusm:    "Single",
			Pekerjaan:  "Website Programmer",
			Kode:       helpers.RandomString(6),
			JoinDate:   "2025-02-14",
			IsVerified: true,
			IsActive:   true,
		}

		// Struct buat data pengguna
		simPengguna := models.User{
			RoleId:     2,
			Name:       "Dodo Sapi",
			Email:      "ggsblackwolf@gmail.com",
			Password:   pass1,
			Jenkel:     "Laki-Laki",
			Phone:      "6285814253237",
			Nik:        "3518110710950003",
			Kota:       "Jakarta",
			Provinsi:   "DKI Jakarta",
			Negara:     "Indonesia",
			Statusm:    "Single",
			Pekerjaan:  "Website Programmer",
			Kode:       helpers.RandomString(6),
			JoinDate:   "2025-02-14",
			IsVerified: true,
			IsActive:   true,
		}

		// Simpan struct data kedalam database
		settings.DB.Create(&simAdm)
		settings.DB.Create(&simPengguna)

		// Tampilkan Printline Seed Sukses
		log.Println("Seeder Pengguna Aplikasi Sukses Dibuat!")
	}
}
