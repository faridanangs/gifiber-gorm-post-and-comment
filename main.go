package main

import (
	"post_comment_api/apps"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	apps.Run(app)
}
