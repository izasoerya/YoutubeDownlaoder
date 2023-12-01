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
	executeDownload(url)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": url,
	})
}

func executeDownload(url string) {
	dir := "C:/Users/ihzas/OneDrive/Dokumen/Coding/Golang/ytdownload/controller"
	interpreterPath := "python3"
	fileName := "download.py"

	cmd := exec.Command(interpreterPath, fileName, url)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
