package controllers

import (
	"fmt"

	model "github.com/DucGiDay/go-fiber-restapi-firebase/models"
	fiber "github.com/gofiber/fiber/v2"
)

func ListCauHoiKep(c *fiber.Ctx) error {
	cauHois, err := model.ListCauHoiKep()
	fmt.Println("debug 1")
	if err != nil {
		return err
	}

	return c.JSON(cauHois)
}
