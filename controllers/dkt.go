package controllers

import (
	"encoding/json"
	// "fmt"

	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func List(c *fiber.Ctx) error {
	dangKienThucs, IDs, err := models.List()
	if err != nil {
		return err
	}
	// respronseDatas := []map[string]interface{}{}
	responseDatas := []string{}
	for i, dangKienThuc := range dangKienThucs {
		dangKienThucByte, _ := json.Marshal(dangKienThuc)
		dangKienThucString := string(dangKienThucByte)
		responseDataString := dangKienThucString + " - id: " + IDs[i]
		responseDatas = append(responseDatas, responseDataString)
	}

	return c.JSON(responseDatas)
}

func Read(c *fiber.Ctx) error {
	id := c.Params("id")
	dangKienThuc, err, _ := models.Read(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(dangKienThuc)
	// return c.JSON(dangKienThucJsonStr)
}

func Create(c *fiber.Ctx) error {
	var dangKienThuc models.DangKienThuc
	if err := c.BodyParser(&dangKienThuc); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Message,
				"error":   err,
			})
		}
	}
	dangKienThuc, err := models.Create(dangKienThuc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON(dangKienThuc)
}

func Update(c *fiber.Ctx) error {
	var dangKienThuc models.DangKienThuc
	id := c.Params("id")
	if err := c.BodyParser(&dangKienThuc); err != nil {
		return err
	}

	dangKienThuc, err := models.Update(id, dangKienThuc)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(dangKienThuc)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON([]byte(`{"message": "Delete Successfuly!"}`))
}
