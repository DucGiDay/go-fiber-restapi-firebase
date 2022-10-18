package route

import (
	controller "github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
	"github.com/gofiber/fiber/v2"
)

func CauHoiRoute(route fiber.Router) {
	route.Get("/", controller.ListCauHoi)
	route.Post("/", controller.CreateCauHoi)
	route.Get("/:id", controller.ReadCauHoi)
	route.Put("/:id", controller.UpdateCauHoi)
	route.Delete("/:id", controller.DeleteCauHoi)
}

func CauHoiKepRoute(route fiber.Router) {
	route.Get("/", controller.ListCauHoiKep)
}

func BothCauHoiRoute(route fiber.Router) {
	route.Get("/", controller.ListBothCauHoi)
}
