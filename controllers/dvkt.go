package controllers

import (

	// "log"

	"encoding/json"

	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func ListDonViKienThucs(c *fiber.Ctx) error {
	donViKienThucs, IDs, err := models.ListDonViKienThucs()
	if err != nil {
		return err
	}
	responseDatas := []string{}
	for i, donViKienThuc := range donViKienThucs {
		donViKienThucByte, _ := json.Marshal(donViKienThuc)
		donViKienThucString := string(donViKienThucByte)
		responseDataString := donViKienThucString + " - id: " + IDs[i]
		responseDatas = append(responseDatas, responseDataString)
	}

	return c.JSON(responseDatas)
}

func ReadDonViKienThuc(c *fiber.Ctx) error {
	// id, err := primitive.ObjectIDFromHex(c.Params("id"))
	id := c.Params("id")
	// if err != nil {
	// 	fmt.Println(3)
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": fiber.ErrBadRequest.Message,
	// 		"error":   err,
	// 	})
	// }
	donViKienThuc, err := models.ReadDonViKienThuc(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(donViKienThuc)
}

func CreateDonViKienThuc(c *fiber.Ctx) error {
	var donViKienThuc models.DonViKienThuc
	if err := c.BodyParser(&donViKienThuc); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Message,
				"error":   err,
			})
		}
	}
	// fmt.Println("donViKienThuc", donViKienThuc)
	// return c.JSON(donViKienThuc)
	donViKienThuc, err := models.CreateDonViKienThuc(donViKienThuc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON(donViKienThuc)
}

func UpdateDonViKienThuc(c *fiber.Ctx) error {
	var donViKienThuc models.DonViKienThuc
	id := c.Params("id")
	// id, err := primitive.ObjectIDFromHex(c.Params("id"))
	// if err != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": fiber.ErrNotFound.Message,
	// 		"error":   err,
	// 	})
	// }
	if err := c.BodyParser(&donViKienThuc); err != nil {
		return err
	}

	donViKienThuc, err := models.UpdateDonViKienThuc(id, donViKienThuc)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(donViKienThuc)
}

func DeleteDonViKienThuc(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteDonViKienThuc(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON([]byte(`{"message": "Delete Successfuly!"}`))
}
