package controllers

import (
	// "fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userId, err := primitive.ObjectIDFromHex(c.Params("userId"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	user, _ := models.GetUser(userId.String())

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Message,
				"error":   err,
			})
		}
	}

	// user.ID = primitive.NewObjectID()
	user.ID = uuid.New()
	user.UserID = user.ID.String()
	user, err := models.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	userId, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
			"error":   err,
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user, err = models.UpdateUser(userId.String(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	userId, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
			"error":   err,
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err = models.DeleteUser(userId.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}
	return c.JSON(user)
}
