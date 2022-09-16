package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
)

func UserRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllUsers)
	route.Post("/", controllers.CreateUser)
	route.Get("/:userId", controllers.GetUser)
	route.Put("/:userId", controllers.UpdateUser)
	route.Delete("/:userId", controllers.DeleteUser)
}
