package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/controller"
)

func UserRoute(route fiber.Router) {
	route.Get("/", controller.GetAllUsers)
	route.Post("/", controller.CreateUser)
	route.Get("/:userId", controller.GetUser)
	route.Put("/:userId", controller.UpdateUser)
}
