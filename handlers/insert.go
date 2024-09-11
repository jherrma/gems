package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/models"
	"github.com/jherrma/gems/services"
)

func InsertGem(mongoService *services.MongoDb) func(ctx fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var gem models.Gem
		if err := c.Bind().Body(&gem); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err := mongoService.InsertGem(&gem)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"gem": gem,
		})
	}
}
