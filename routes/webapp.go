package routes

import (
	"log"
	"net/http"
	"smartsaver/controllers"
	"smartsaver/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebAppRoute() {
	// Memanggil fungsi route dari framework gin golang
	router := gin.Default()

	// Menambahkan cors pada settingan route gin golang
	router.Use(cors.Default())

	// Menggunakan middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://example.com"}, // Ganti dengan origin yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Memuat template dengan ekstensi .tmpl dari direktori view
	router.LoadHTMLGlob("views/*.tmpl")

	// Menyajikan file statis dari direktori assets
	router.Static("/assets", "./assets")

	// Route Website APP
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Login.tmpl", nil) // Render template Login.tmpl
	})
	router.GET("/halaman-autentikasi/pendaftaran-akun", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Register.tmpl", nil) // Render template Register.tmpl
	})
	router.GET("/halaman-aplikasi/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Dashboard.tmpl", nil) // Render template Dashboard.tmpl
	})
	router.GET("/halaman-aplikasi/profile-pengguna", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Profile.tmpl", nil) // Render template Dashboard.tmpl
	})

	// Route API
	v1 := router.Group("api/v1/")
	{
		// Route untuk aplikasi
		oth := v1.Group("/autentikasi/")
		{
			oth.POST("/login", controllers.CheckLogin)
			oth.POST("/register", controllers.RegisterAkun)
			oth.GET("/verifikasi/:kode", controllers.VerifikasiAkun)
		}

		apps := v1.Group("/main-app/")
		{
			dashApps := apps.Group("/dashboard/")
			{
				dashApps.GET("/check-token", middleware.IsAuth(), controllers.CheckToken)
				dashApps.POST("/logout", controllers.Logout)
			}
		}
	}

	// Menampilkan log server berjalan dengan port 8080
	log.Println("Server started on: http://127.0.0.1:8080")

	// Menjalankan server ke port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
