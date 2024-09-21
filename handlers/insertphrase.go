package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/models"
	"github.com/jherrma/gems/services"
)

func InsertPhrase(mongoService *services.MongoDb) func(ctx fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var phraseRequest models.PhraseRequest
		if err := c.Bind().Body(&phraseRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		gem := services.ComputeGem(phraseRequest.Phrase)

		err := mongoService.InsertGem(gem)
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
