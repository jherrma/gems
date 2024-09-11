package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/services"
)

func GetList(mongoService *services.MongoDb) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		skipString := c.Query("skip")
		limitString := c.Query("limit")

		skip, err := strconv.ParseInt(skipString, 10, 64)
		if err != nil || skip < 0 {
			skip = 0
		}

		limit, err := strconv.ParseInt(limitString, 10, 64)
		if err != nil || limit <= 0 {
			limit = 10
		}

		gems, err := mongoService.GetGems(skip, limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": gems,
		})
	}
}
