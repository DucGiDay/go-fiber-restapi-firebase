package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	model "github.com/DucGiDay/go-fiber-restapi-firebase/models"
	fiber "github.com/gofiber/fiber/v2"
)

func ListCauHoi(c *fiber.Ctx) error {
	cauHois, Ids, err := model.ListCauHoi()
	if err != nil {
		return err
	}
	responseDatas := []string{}
	for i, cauhoi := range cauHois {
		cauhoiByte, _ := json.Marshal(cauhoi)
		cauhoiString := string(cauhoiByte)
		fmt.Println(cauhoiString)
		responseDataString := cauhoiString + " - id: " + Ids[i]
		responseDatas = append(responseDatas, responseDataString)
	}

	return c.JSON(responseDatas)
}

func ReadCauHoi(c *fiber.Ctx) error {
	id := c.Params("id")
	cauHoi, err := model.ReadCauHoi(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(cauHoi)
}

func CreateCauHoi(c *fiber.Ctx) error {
	var cauHoi model.Cauhoi
	if err := c.BodyParser(&cauHoi); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"succes":  false,
				"message": fiber.ErrBadRequest.Message,
				"error":   err,
			})
		}
	}
	cauHoi.Date = time.Now()
	cauHoi, err := model.CreateCauHoi(cauHoi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"succes":  false,
			"message": fiber.ErrInternalServerError.Message,
			"erroe":   err,
		})
	}
	return c.JSON(cauHoi)
}

func UpdateCauHoi(c *fiber.Ctx) error {
	var cauHoi model.Cauhoi
	id := c.Params("id")
	if err := c.BodyParser(&cauHoi); err != nil {
		return err
	}

	cauHoi, err := model.UpdateCauHoi(id, cauHoi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON(cauHoi)
}

func DeleteCauHoi(c *fiber.Ctx) error {
	id := c.Params("id")
	err := model.DeleteCauHoi(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"succes":  false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON([]byte(`{"message"}:"Delete Successfuly!"`))
}
