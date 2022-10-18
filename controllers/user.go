package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"encoding/json"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, UId, err := models.GetAllUsers()
	if err != nil {
		return err
	}
	responseDatas := []string{}
	for i, cauhoi := range users {
		UserByte, _ := json.Marshal(cauhoi)
		UserString := string(UserByte)
		fmt.Println(UserString)
		responseDataString := UserString + " - id: " + UId[i]
		responseDatas = append(responseDatas, responseDataString)
	}
	return c.JSON(responseDatas)
}
func GetAllAdmins(c *fiber.Ctx) error {
	users, UId, err := models.GetAllAdmins()
	if err != nil {
		return err
	}
	responseDatas := []string{}
	for i, cauhoi := range users {
		UserByte, _ := json.Marshal(cauhoi)
		UserString := string(UserByte)
		fmt.Println(UserString)
		responseDataString := UserString + " - id: " + UId[i]
		responseDatas = append(responseDatas, responseDataString)
	}
	return c.JSON(responseDatas)
}

func GetUser(c *fiber.Ctx) error {
	userId:= c.Params("userId")
	user, err := models.GetUser(userId)
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
	userId := c.Params("userId")
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user, err := models.UpdateUser(userId, user)

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
