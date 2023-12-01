package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/ytdownload/controller"
)

func SetupRouter(route fiber.Router) {
	route.Post("", controller.ParseLinkYoutube)
}
