package helpers

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"smartsaver/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func CreateToken(user *models.User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"role_id":     user.RoleId,
		"name":        user.Name,
		"email":       user.Email,
		"jenkel":      user.Jenkel,
		"phone":       user.Phone,
		"nik":         user.Nik,
		"kota":        user.Kota,
		"provinsi":    user.Provinsi,
		"negara":      user.Negara,
		"statusm":     user.Statusm,
		"pekerjaan":   user.Pekerjaan,
		"kode":        user.Kode,
		"join_date":   user.JoinDate,
		"is_verified": user.IsVerified,
		"is_active":   user.IsActive,
		"exp":         time.Now().AddDate(0, 0, 7).Unix(),
		"iat":         time.Now().Unix(),
	})

	// Sign and get the completed encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

// Fungsi untuk mengirim email
func SendRegisEmail(to string, subject string, body string) error {
	from := "mrvisca2018@gmail.com"   // Ganti dengan email Anda
	password := "pckl lpst nlqb snal" // Ganti dengan password email Anda

	// Set up the SMTP server configuration
	smtpHost := "smtp.gmail.com" // Ganti dengan SMTP host Anda
	smtpPort := "587"            // Port SMTP (umumnya 587 untuk TLS)

	// Buat message
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Kirim email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}

// Fungsi untuk membuat string acak 6 karakter
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Buat random generator baru dengan seed berbasis waktu
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, n)
	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}
