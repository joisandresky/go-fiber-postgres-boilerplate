package v1

import "github.com/gofiber/fiber/v2"

func (api *httpApi) RegisterRoutes(app fiber.Router) {
	exampleApi := app.Group("/api/v1/examples")

	exampleApi.Post("/say-hello", api.SayHello)
}
