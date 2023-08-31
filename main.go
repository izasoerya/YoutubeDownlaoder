package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/izasoerya/ytdownload/router"
)

func main() {
	/* Generate static HTML file */
	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	/* Serve HTML as index in /home */
	app.Get("/home", func(c *fiber.Ctx) error {						//* set index.html in /app
		return c.Render("index", fiber.Map{})
	})
	app.Static("/static", "./static")



	/* Logger middleware */
	app.Use(logger.New())

	/* Give router endpoint */
	router.SetupRouter(app.Group("/download"))	

	if err:= app.Listen(":3000"); err != nil {
		panic(err);
	}
}
