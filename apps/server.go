package apps

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Run(app *fiber.App) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	db := InitializeDB(dbConfig)
	InitializeRouter(app, db)
	app.Listen(":8000")
}
