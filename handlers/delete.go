package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/losevs/people-api/database"
	"github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
)

func DeleteByID(c *fiber.Ctx) error {
	Person := new(models.PersonResponse)
	needID := c.Params("id", "0")
	if needID == "0" {
		logger.Logg.Debug("Parsing params DelById error")
		return c.Status(fiber.StatusInternalServerError).JSON("Params error. Try Again.")
	}
	needIdInt, err := strconv.Atoi(needID)
	if err != nil {
		logger.Logg.Debug("strconv error with ID: ", err)
		return c.Status(fiber.StatusBadRequest).JSON("Wrong ID input.")
	}
	if checker := database.DB.Db.Where("id = ?", needIdInt).First(Person).Delete(Person); checker.RowsAffected == 0 {
		logger.Logg.Info(fmt.Sprint("No person with id = ", needIdInt))
		return c.Status(fiber.StatusBadRequest).JSON(fmt.Sprint("There is no person with id = ", needIdInt))
	}
	logger.Logg.Info(fmt.Sprintf("Person with id=%d has been deleted", needIdInt))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "person successfully deleted",
		"person":  *Person,
	})
}
