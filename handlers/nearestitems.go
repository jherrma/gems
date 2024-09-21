package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/models"
	"github.com/jherrma/gems/services"
)

func GetNearesItems(mongoService *services.MongoDb) func(ctx fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var distanceRequest models.NearestItemsRequest
		if err := c.Bind().Body(&distanceRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		nearesPhrases, err := services.GetNearestItemsToPhrase(distanceRequest.Phrase, distanceRequest.Limit, mongoService)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		referenceGem, err := mongoService.GetGem(distanceRequest.Phrase)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"reference":       referenceGem,
			"nearest_phrases": nearesPhrases,
		})
	}
}
