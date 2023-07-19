package controller

import (
	"fmt"
	"os"
	"os/exec"

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
	fetchPython(url)
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"OK",
		"message": url,
	})
}

func fetchPython(url string) {
	dir := "/home/iza/Documents/Coding/Golang/src/downloaderYT/python"
	interpreterPath := "/home/iza/Documents/Coding/Golang/src/downloaderYT/python/venv/bin/python3"
	fileName := "test.py"

	cmd := exec.Command(interpreterPath, fileName, url)
	cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
