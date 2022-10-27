package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
)

func AdminRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllAdmins)
	route.Post("/", controllers.CreateUser)
	route.Get("/:userId", controllers.GetUser)
	route.Put("/:userId", controllers.UpdateUser)
	route.Delete("/:userId", controllers.DeleteUser)
}
