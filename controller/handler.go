package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserInputURL(c *fiber.Ctx) error{
	type Request struct {
		URL string `json:"url"`
	}
	var body Request
	err := c.BodyParser(&body)
	if err!= nil {
        fmt.Println(err)
    }

	url := body.URL
	fmt.Println(url)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"OK",
		"message": url,
	})
}
