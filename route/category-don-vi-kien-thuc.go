package route

import (
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
	"github.com/gofiber/fiber/v2"
)

func DonViKienThucRoute(route fiber.Router) {
	route.Get("/", controllers.ListDonViKienThucs)
	route.Post("/", controllers.CreateDonViKienThuc)
	route.Get("/:id", controllers.ReadDonViKienThuc)
	route.Put("/:id", controllers.UpdateDonViKienThuc)
	route.Delete("/:id", controllers.DeleteDonViKienThuc)
}
