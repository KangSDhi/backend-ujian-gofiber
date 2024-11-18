package routers

import (
	"backend-ujian-gofiber/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/ping", controllers.Ping)

	apiAuth := api.Group("/auth")

	apiAuth.Post("/login", controllers.Login)
}
