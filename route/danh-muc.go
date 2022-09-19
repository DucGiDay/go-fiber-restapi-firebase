package route

import (
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
	"github.com/gofiber/fiber/v2"
)

func DanhMucRoute(route fiber.Router) {
	route.Get("/", controllers.ListDanhMuc)
}
