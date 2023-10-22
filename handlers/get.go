package handlers

import (
	"fmt"
	"strconv"

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
	Person := models.PersonResponse{}
	needID := c.Params("id", "0")
	if needID == "0" {
		logger.Logg.Debug("Parsing params getById error")
		return c.Status(fiber.StatusInternalServerError).JSON("Params error. Try Again.")
	}
	needIdInt, err := strconv.Atoi(needID)
	if err != nil {
		logger.Logg.Debug("strconv error with ID: ", err)
		return c.Status(fiber.StatusBadRequest).JSON("Wrong ID input.")
	}
	if checker := database.DB.Db.Where("id = ?", needIdInt).First(&Person); checker.RowsAffected == 0 {
		logger.Logg.Info("No person with id = ", needIdInt)
		return c.Status(fiber.StatusOK).JSON(fmt.Sprint("There is no person with id = ", needIdInt))
	}
	return c.Status(fiber.StatusOK).JSON(Person)
}
