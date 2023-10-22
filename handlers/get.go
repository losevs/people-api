package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/losevs/people-api/database"
	"github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
)

// Вывод всех
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

// Вывод по ID
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

// Пагинация
func Pag(c *fiber.Ctx) error {
	needPage, err := strconv.Atoi(c.Params("page", "1"))
	if err != nil {
		logger.Logg.Debug("page parsing error", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	People := []models.PersonResponse{}
	if check := database.DB.Db.Offset((needPage - 1) * 5).Limit(5).Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("Page is empty")
		return c.Status(fiber.StatusOK).JSON("This page is empty")
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// Пагинация по м. полу
func MenPag(c *fiber.Ctx) error {
	needPage, err := strconv.Atoi(c.Params("page", "1"))
	if err != nil {
		logger.Logg.Debug("page parsing error", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	People := []models.PersonResponse{}
	if check := database.DB.Db.Offset((needPage-1)*4).Limit(4).Where("sex = ?", "male").Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("There is no males in the db")
		return c.Status(fiber.StatusOK).JSON("DB is empty")
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// Пагинация по ж. полу
func WMenPag(c *fiber.Ctx) error {
	needPage, err := strconv.Atoi(c.Params("page", "1"))
	if err != nil {
		logger.Logg.Debug("page parsing error", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	People := []models.PersonResponse{}
	if check := database.DB.Db.Offset((needPage-1)*4).Limit(4).Where("sex = ?", "female").Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("There is no males in the db")
		return c.Status(fiber.StatusOK).JSON("DB is empty")
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// Возраст - по возрастанию
func AgeType(c *fiber.Ctx) error {
	People := []models.PersonResponse{}
	if check := database.DB.Db.Order("age asc").Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("DB is empty - ageType")
		return c.Status(fiber.StatusOK).JSON("DB is empty")
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// Фильтрация
// По возрасту
func FiltAge(c *fiber.Ctx) error {
	needAge, err := strconv.Atoi(c.Params("age"))
	if err != nil {
		logger.Logg.Debug(err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	People := []models.PersonResponse{}
	if check := database.DB.Db.Where("age = ?", needAge).Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("db is empty with age = ", needAge)
		return c.Status(fiber.StatusBadRequest).JSON(fmt.Sprint("There is no people with age = ", needAge))
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// По стране
func FiltCountry(c *fiber.Ctx) error {
	needCountry := c.Params("country")
	needCountry = strings.ToUpper(needCountry)
	People := []models.PersonResponse{}
	if check := database.DB.Db.Where("nationality = ?", needCountry).Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("db is empty with country = ", needCountry)
		return c.Status(fiber.StatusBadRequest).JSON(fmt.Sprint("There is no people from ", needCountry))
	}
	return c.Status(fiber.StatusOK).JSON(People)
}

// По полу
func FiltSex(c *fiber.Ctx) error {
	needGen := c.Params("sex")
	People := []models.PersonResponse{}
	if check := database.DB.Db.Where("sex = ?", needGen).Find(&People); check.RowsAffected == 0 {
		logger.Logg.Info("db is empty with gen = ", needGen)
		return c.Status(fiber.StatusBadRequest).JSON(fmt.Sprintf("There is %s people", needGen))
	}
	return c.Status(fiber.StatusOK).JSON(People)
}
