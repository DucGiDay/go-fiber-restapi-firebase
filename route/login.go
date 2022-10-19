package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DucGiDay/go-fiber-restapi-firebase/controllers"
)

func LoginRoute(route fiber.Router) {
	route.Post("/", controllers.Login)
}
