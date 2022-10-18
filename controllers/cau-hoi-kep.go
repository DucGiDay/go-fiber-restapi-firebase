package controllers

import (
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
