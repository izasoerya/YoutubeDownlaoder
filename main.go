package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/izasoerya/ytdownload/router"
)

func setupRoutes(app *fiber.App) {									//* Create home page
	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "api endpoint",
		})
	})
	router.SetupRouter(api.Group("/link"))
}

func main() {
	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	app.Use(logger.New())
	setupRoutes(app)
	
	app.Get("/home", func(c *fiber.Ctx) error {						//* set index.html in /app
		return c.Render("index", fiber.Map{})
	})

	

	app.Listen(":3000")
}
