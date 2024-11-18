package main

import (
	"backend-ujian-gofiber/src/database"
	"backend-ujian-gofiber/src/models"
	"backend-ujian-gofiber/src/routers"
	"backend-ujian-gofiber/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.InitDB()
	database.DB.AutoMigrate(&models.Kelas{})
	database.DB.AutoMigrate(&models.Pengguna{})
	seedData()
}

func seedData() {
	emailAdmin := "kangsigit@gmail.com"
	passwordAdmin := "qwerty"
	passwordAdminHash, _ := utils.HashPassword(passwordAdmin)

	idSiswa := "X-TKJ-1"
	passwordSiswa := "ytrewq"
	passwordSiswaHash, _ := utils.HashPassword(passwordSiswa)

	var admin = []models.Pengguna{
		{
			ID:            uuid.New(),
			IdSiswa:       nil,
			NamaPengguna:  "Sigit Admin",
			EmailPengguna: &emailAdmin,
			Password:      passwordAdminHash,
			PasswordPlain: nil,
			RolePengguna:  models.Admin,
			KelasID:       nil,
		},
	}

	var siswa = []models.Pengguna{
		{
			ID:            uuid.New(),
			IdSiswa:       &idSiswa,
			NamaPengguna:  "Sigit Siswa",
			EmailPengguna: nil,
			Password:      passwordSiswaHash,
			PasswordPlain: &passwordSiswa,
			RolePengguna:  models.Siswa,
			KelasID:       nil,
		},
	}

	database.DB.Create(&admin)
	database.DB.Create(&siswa)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Successfully loaded .env file")
}

func serveApplication() {
	app := fiber.New()

	routers.SetupRoutes(app)

	err := app.Listen(":8081")
	if err != nil {
		log.Fatal(err)
		return
	}
}
