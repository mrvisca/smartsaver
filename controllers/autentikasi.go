package controllers

import (
	"errors"
	"net/http"
	"smartsaver/helpers"
	"smartsaver/models"
	"smartsaver/settings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FillResRegister(moduser models.User) models.ResRegister {
	return models.ResRegister{
		ID:         moduser.ID,
		RoleId:     moduser.ID,
		Name:       moduser.Name,
		Email:      moduser.Email,
		Jenkel:     moduser.Jenkel,
		Phone:      moduser.Phone,
		Nik:        moduser.Nik,
		Kota:       moduser.Kota,
		Provinsi:   moduser.Provinsi,
		Negara:     moduser.Negara,
		Statusm:    moduser.Statusm,
		Pekerjaan:  moduser.Pekerjaan,
		Kode:       moduser.Kode,
		JoinDate:   moduser.JoinDate,
		IsVerified: moduser.IsVerified,
		IsActive:   moduser.IsActive,
	}
}

func CheckLogin(c *gin.Context) {
	// Definisi body request kedalam variabel
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Variabel model user
	var usr models.User

	// Kondisi bila data by email tidak ditemukan
	if errors.Is(settings.DB.Where("email LIKE ? AND is_active = ?", "%"+email+"%", true).First(&usr).Error, gorm.ErrRecordNotFound) {
		helpers.ElorResponse(c, "Data pengguna tidak ditemukan dalam sistem kami!")
		c.Abort()
		return
	}

	// Kondisi jika akun belum terverifikasi
	if !usr.IsVerified {
		helpers.ElorResponse(c, "Akun belum melakukan verifikasi email!")
		c.Abort()
		return
	}

	// Kondisi jika memasukan password salah
	if !helpers.CheckPassword(usr.Password, password) {
		helpers.ElorResponse(c, "Password tidak valid!")
		c.Abort()
		return
	}

	// Buat token bila data sudah sesuai dengan data dalam database
	token := helpers.CreateToken(&usr)

	// Tampilkan response sukses
	helpers.TokenSuksesResponse(c, "Login akun pengguna berhasil dilakukan!", token)
}

func RegisterAkun(c *gin.Context) {
	// Definisikan body request kedalam variabel
	name := c.PostForm("name")
	email := c.PostForm("email")
	pass := c.PostForm("password")
	phone := c.PostForm("phone")
	jenkel := c.PostForm("jenkel")

	// variabel model user
	var usr models.User

	// Kondisi pencegahan duplikasi data berdasarkan nama dan email
	if !errors.Is(settings.DB.Where("email LIKE ? AND name LIKE ?", "%"+email+"%", "%"+name+"%").First(&usr).Error, gorm.ErrRecordNotFound) {
		helpers.ElorResponse(c, "Email atau nama akun sudah terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Enkripsi Password
	pw, _ := helpers.HashPassword(pass)

	// Tanggal bergabung otomatis
	sekarang := time.Now()
	join := sekarang.Format("2006-01-02 15:04:05")

	// Simpan struct data penggun baru
	simpan := models.User{
		RoleId:     2,
		Name:       name,
		Email:      email,
		Password:   pw,
		Jenkel:     jenkel,
		Phone:      phone,
		Kode:       helpers.RandomString(6),
		JoinDate:   join,
		IsVerified: false,
		IsActive:   true,
	}

	// Simpan struct data ke database
	settings.DB.Create(&simpan)

	// Deklarasi variabel untuk kebutuhan email verifikasi akun baru pengguna
	link := "http://localhost:8080/api/v1/autentikasi/verifikasi/" + simpan.Kode
	to := simpan.Email
	subject := "Email verifikasi akun baru!"
	body := "Klik link ini untuk melakukan verifikasi akun baru! \n" + link

	// Kirim email verifikasi akun baru ke email pengguna yang sudah terdaftar
	err := helpers.SendRegisEmail(to, subject, body)
	if err != nil {
		helpers.ElorResponse(c, err.Error())
		return
	}

	// Struct response data
	res := FillResRegister(simpan)

	// Tampilkan helper response sukses
	helpers.SuksesWithData(c, "Berhasil mendaftarkan akun baru, jangan lupa untuk verifikasi email akun anda!", res)
}

func VerifikasiAkun(c *gin.Context) {
	// Definisikan param kode kedalam variabel
	kode := c.Param("kode")

	// Variabel model user
	var usr models.User

	// Kondisi bila kode pengguna tidak ditemukan
	if errors.Is(settings.DB.Where("kode LIKE ? AND is_active = ?", "%"+kode+"%", true).First(&usr).Error, gorm.ErrRecordNotFound) {
		helpers.ElorResponse(c, "Kode tidak valid / terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Kondisi bila akun telah terverifikasi sebelumnya
	if usr.IsVerified {
		helpers.ElorResponse(c, "Akun sudah diverifikasi sebelumnya!")
		c.Abort()
		return
	}

	// Verifikasi akun pengguna
	settings.DB.Model(&usr).Where("kode = ?", kode).Update("is_verified", true)

	// Redirect ke halaman login
	c.Redirect(http.StatusFound, "http://localhost:8080/")
}
