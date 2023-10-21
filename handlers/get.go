package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/losevs/people-api/database"
	"github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
)

func ShowAll(c *fiber.Ctx) error {
	People := []models.PersonResponse{}
	checker := database.DB.Db.Find(&People)
	if checker.Error != nil {
		logger.Logg.Debug("Error database find ShowAll func: ", checker.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(checker.Error.Error())
	}
	if checker.RowsAffected == 0 {
		logger.Logg.Info("Database is empty")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "database is empty",
		})
	}
	return c.JSON(People)
}

func ShowByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "everything is ok!",
	})
}
