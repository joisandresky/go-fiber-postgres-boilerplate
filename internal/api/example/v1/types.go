package v1

import "github.com/gofiber/fiber/v2"

type HttpAPI interface {
	RegisterRoutes(app fiber.Router)
}
