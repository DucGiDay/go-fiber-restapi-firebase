package controllers

import (
	"time"

	model "github.com/DucGiDay/go-fiber-restapi-firebase/models"
	fiber "github.com/gofiber/fiber/v2"
)

func ListCauHoiKep(c *fiber.Ctx) error {
	cauHois, err := model.ListCauHoiKep()
	if err != nil {
		return err
	}

	return c.JSON(cauHois)
}
func CreateCauHoiKep(c *fiber.Ctx) error {
	var cauHoi model.CauHoiKep
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
	cauHoi.Don_Kep = 1
	cauHoi, err := model.CreateCauHoiKep(cauHoi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"succes":  false,
			"message": fiber.ErrInternalServerError.Message,
			"erroe":   err,
		})
	}
	return c.JSON(cauHoi)
}
