package handlers

import "github.com/gofiber/fiber/v2"

func DeleteByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "delete is ok!",
	})
}