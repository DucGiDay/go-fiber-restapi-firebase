package route

import (
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
	"github.com/gofiber/fiber/v2"
)

func MoTaChiTietRoute(route fiber.Router) {
	route.Get("/", controllers.ListMoTaChiTiets)
	route.Post("/", controllers.CreateMoTaChiTiet)
	route.Get("/:id", controllers.ReadMoTaChiTiet)
	route.Put("/:id", controllers.UpdateMoTaChiTiet)
	route.Delete("/:id", controllers.DeleteMoTaChiTiet)
}
