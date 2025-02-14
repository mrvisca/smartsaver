package main

import (
	"smartsaver/routes"
	"smartsaver/seeders"
	"smartsaver/settings"
)

func main() {
	// Panggil fungsi koneksi ke database
	settings.InitDB()
	defer settings.DB.Close()

	seeders.RoleSeed(nil)
	seeders.UserSeed(nil)

	// Panggil fungsi route webapp
	routes.WebAppRoute()
}
