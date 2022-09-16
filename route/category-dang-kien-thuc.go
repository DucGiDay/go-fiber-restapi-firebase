package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
)

func DangKienThucRoute(route fiber.Router) {
	route.Get("/", controllers.List)
	route.Post("/", controllers.Create)
	route.Get("/:id", controllers.Read)
	route.Put("/:id", controllers.Update)
	route.Delete("/:id", controllers.Delete)
}
