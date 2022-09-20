package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
)

func CauHoiRoute(route fiber.Router) {
	route.Get("/", controller.ListCauHoi)
	route.Post("/", controller.CreateCauHoi)
	route.Get("/:id", controller.ReadCauHoi)
	route.Put("/:id", controller.UpdateCauHoi)
	route.Delete("/:id", controller.DeleteCauHoi)
}