package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"smartsaver/helpers"
	"smartsaver/models"
	"smartsaver/settings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FillProfile(usr models.User) models.ProfilePengguna {
	joinDateFormatted := "" // Inisialisasi string kosong

	if usr.JoinDate != "" { // Pastikan JoinDate tidak kosong
		// 1. Parsing string ke time.Time
		t, err := time.Parse("2006-01-02 15:04:05", usr.JoinDate) // Format input yang sesuai
		if err != nil {
			fmt.Println("Error parsing JoinDate:", err)
			// Tangani error, misalnya dengan memberikan nilai default atau log error
			// joinDateFormatted akan tetap "" jika error
		} else {
			// 2. Format time.Time ke string YYYY-MM-DD
			joinDateFormatted = t.Format("2006-01-02")
		}
	}

	return models.ProfilePengguna{
		ID:         usr.ID,
		RoleId:     usr.RoleId,
		RoleName:   usr.Role.Name,
		Name:       usr.Name,
		Email:      usr.Email,
		Jenkel:     usr.Jenkel,
		Phone:      usr.Phone,
		Kota:       usr.Kota,
		Provinsi:   usr.Provinsi,
		Negara:     usr.Negara,
		Statusm:    usr.Statusm,
		Pekerjaan:  usr.Pekerjaan,
		Kode:       usr.Kode,
		JoinDate:   joinDateFormatted,
		IsVerified: usr.IsVerified,
		IsActive:   usr.IsActive,
	}
}

func CheckToken(c *gin.Context) {
	// Definisi data token
	userid := uint(c.MustGet("jwt_user_id").(float64))
	roleid := uint(c.MustGet("jwt_role_id").(float64))
	name := c.MustGet("jwt_name").(string)
	email := c.MustGet("jwt_email").(string)

	// Definisi model user
	var moduser models.User

	// Kondisi validasi data
	if errors.Is(settings.DB.Where("id = ? AND role_id = ? AND name LIKE ? AND email LIKE ?", userid, roleid, "%"+name+"%", "%"+email+"%").Preload("Role").First(&moduser).Error, gorm.ErrRecordNotFound) {
		helpers.ElorResponse(c, "Terjadi kesalahan saat validasi data token!")
		c.Abort()
		return
	}

	// Isi struct response dengan data
	response := FillProfile(moduser)

	// Tampilkan response
	helpers.SuksesWithData(c, "Berhasil memanggil data profile!", response)
}

// Fungsi untuk melakukan logout
func Logout(c *gin.Context) {
	// Hapus cookie JWT (sesuaikan dengan nama cookie Anda)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Menghapus cookie
	})

	helpers.SuksesResponse(c, "Berhasil melakukan logout akun!")
}
