package controllers

import (

	// "log"

	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func ListMoTaChiTiets(c *fiber.Ctx) error {
	moTaChiTiets, err := models.ListMoTaChiTiets()
	if err != nil {
		return err
	}

	return c.JSON(moTaChiTiets)
}

func ReadMoTaChiTiet(c *fiber.Ctx) error {
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
	moTaChiTiet, err := models.ReadMoTaChiTiet(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(moTaChiTiet)
}

func CreateMoTaChiTiet(c *fiber.Ctx) error {
	var moTaChiTiet models.MoTaChiTiet
	if err := c.BodyParser(&moTaChiTiet); err != nil {
		return err
	}

	moTaChiTiet, err := models.CreateMoTaChiTiet(moTaChiTiet)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(moTaChiTiet)
}

func UpdateMoTaChiTiet(c *fiber.Ctx) error {
	var moTaChiTiet models.MoTaChiTiet
	id := c.Params("id")
	// id, err := primitive.ObjectIDFromHex(c.Params("id"))
	// if err != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": fiber.ErrNotFound.Message,
	// 		"error":   err,
	// 	})
	// }
	if err := c.BodyParser(&moTaChiTiet); err != nil {
		return err
	}

	moTaChiTiet, err := models.UpdateMoTaChiTiet(id, moTaChiTiet)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(moTaChiTiet)
}

func DeleteMoTaChiTiet(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteMoTaChiTiet(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON([]byte(`{"message": "Delete Successfuly!"}`))
}
