package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/losevs/people-api/database"
	"github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
)

func PatchByID(c *fiber.Ctx) error {
	Person := new(models.PersonResponse)

	needID := c.Params("id", "0")
	if needID == "0" {
		logger.Logg.Debug("Parsing params PatchById error")
		return c.Status(fiber.StatusInternalServerError).JSON("Params error. Try Again.")
	}
	needIdInt, err := strconv.Atoi(needID)
	if err != nil {
		logger.Logg.Debug("strconv error with ID: ", err)
		return c.Status(fiber.StatusBadRequest).JSON("Wrong ID input.")
	}
	if check := database.DB.Db.Where("id = ?", needIdInt).First(Person); check.RowsAffected == 0 {
		logger.Logg.Info(fmt.Sprint("No person with id = ", needIdInt))
		return c.Status(fiber.StatusOK).JSON(fmt.Sprint("There is no person with id = ", needIdInt))
	}

	PersonReq := new(models.PersonRequest)
	if err := c.BodyParser(PersonReq); err != nil {
		logger.Logg.Debug("Error while parsing the body: patch")
		return c.Status(fiber.StatusInternalServerError).JSON("Body parsing error. Try again.")
	}
	if len(PersonReq.Name) != 0 {
		age, gender, country, err := UnMar(Person.Name)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		Person.Name = PersonReq.Name
		Person.Age = age
		Person.Sex = gender
		Person.Nationality = country
	}

	if len(PersonReq.Surname) != 0 {
		Person.Surname = PersonReq.Surname
	}
	if len(PersonReq.Patronymic) != 0 {
		Person.Patronymic = PersonReq.Patronymic
	}
	if check := database.DB.Db.Where("id = ?", needIdInt).Save(Person); check.Error != nil {
		logger.Logg.Debug("Error while saving patched struct")
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	logger.Logg.Info(fmt.Sprintf("Person with id = %d has been patched", needIdInt))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Person successfully patched",
		"person":  Person,
	})
}

/*Patch:
localhost:80/change/2192017524
body:
{
	"name": "qwe",
	"surname": "ewq",
	"patronymic": "asd"
}

response:
models.PersonResponse{...}
*/
