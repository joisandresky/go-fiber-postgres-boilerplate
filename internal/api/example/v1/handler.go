package v1

import (
	"go-fiber-postgres-boilerplate/internal/dto"
	"go-fiber-postgres-boilerplate/internal/usecase/example"

	"github.com/gofiber/fiber/v2"
)

type httpApi struct {
	uc example.Usecase
}

func NewHttpApi(uc example.Usecase) HttpAPI {
	return &httpApi{
		uc,
	}
}

func (api *httpApi) SayHello(c *fiber.Ctx) error {
	req := dto.SayHelloRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to get request body",
		})
	}

	resp, _ := api.uc.SayHello(c.Context(), req)

	return c.Status(resp.Status).JSON(resp)
}
