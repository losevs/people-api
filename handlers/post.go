package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/losevs/people-api/database"
	"github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
)

// Добавление нового человека
func AddNew(c *fiber.Ctx) error {
	persReq := models.PersonRequest{}
	if err := c.BodyParser(&persReq); err != nil {
		logger.Logg.Debug(err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	ageApi, gender, nation, err := UnMar(persReq.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	filledPerson := models.PersonResponse{
		ID:          int64(uuid.New().ID()),
		Name:        persReq.Name,
		Surname:     persReq.Surname,
		Patronymic:  persReq.Patronymic,
		Age:         ageApi,
		Sex:         gender,
		Nationality: nation,
	}
	if checker := database.DB.Db.Create(&filledPerson); checker.Error != nil {
		logger.Logg.Debug("Error creating new person: ", checker.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(checker.Error)
	}
	logger.Logg.Info("New person has been created")
	return c.Status(fiber.StatusOK).JSON(filledPerson)
}

// Запросы к API age, gender, nation
func UnMar(name string) (int64, string, string, error) {
	requestUrlAge := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	requestUrlGender := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	requestUrlNation := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)

	reqAge, err := http.NewRequest(http.MethodGet, requestUrlAge, nil)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	reqGen, err := http.NewRequest(http.MethodGet, requestUrlGender, nil)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	reqNat, err := http.NewRequest(http.MethodGet, requestUrlNation, nil)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}

	resAge, err := http.DefaultClient.Do(reqAge)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	resGen, err := http.DefaultClient.Do(reqGen)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	resNat, err := http.DefaultClient.Do(reqNat)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}

	bodyAge, err := io.ReadAll(resAge.Body)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	bodyGen, err := io.ReadAll(resGen.Body)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}
	bodyNat, err := io.ReadAll(resNat.Body)
	if err != nil {
		logger.Logg.Debug(err)
		return 0, "", "", err
	}

	AgeUnM := models.AgeMarsh{}
	err = json.Unmarshal(bodyAge, &AgeUnM)
	if err != nil {
		logger.Logg.Debug(err)
	}

	GenUnM := models.GenderMarsh{}
	err = json.Unmarshal(bodyGen, &GenUnM)
	if err != nil {
		logger.Logg.Debug(err)
	}

	NatUnM := models.NationMarsh{}
	err = json.Unmarshal(bodyNat, &NatUnM)
	if err != nil {
		logger.Logg.Debug(err)
	}

	if AgeUnM.Age == 0 || GenUnM.Gender == "" || len(NatUnM.Country) == 0 {
		logger.Logg.Info("Strange name input")
		return 0, "", "", errors.New("strange name input - try again")
	}

	return AgeUnM.Age, GenUnM.Gender, NatUnM.Country[0].CountryID, nil
}
