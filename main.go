package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Get("/home", func(c *fiber.Ctx) error {						//* set index.html in /app
		return c.Render("index", fiber.Map{})
	})

	app.Listen(":3000")
}
