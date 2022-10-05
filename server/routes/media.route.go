package routes

import "github.com/gofiber/fiber/v2"

func SetupMediaRoutesV1(router fiber.Router) {
	media := router.Group("/media")

	media.Get("/")
}
