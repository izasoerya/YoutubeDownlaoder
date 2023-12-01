package controller

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func ParseLinkYoutube(c *fiber.Ctx) error {
	type Request struct {
		URL string `json:"url"`
	}
	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
	}
	url := body.URL
	fmt.Println(url)

	errorCode := executeDownload(url)
	if errorCode == -1 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Error at parsing Pytube",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": url,
	})
}

func executeDownload(url string) int8 {
	dir := os.Getenv("DIRECTORY")
	interpreterPath := os.Getenv("INTERPRETER_PATH")
	fileName := "download.py"

	cmd := exec.Command(interpreterPath, fileName, url)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return -1
	}
	return 0
}
