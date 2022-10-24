package main

import (
	"fmt"
	"log"

	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"github.com/DucGiDay/go-fiber-restapi-firebase/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func main() {
	config.ConnectDatabase()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	setupRoutes(app)
	// route.LoginRoute(app)
	defer config.FI.Client.Close()

	// Optional middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if c.Get("host") == "localhost:4000" {
			c.Locals("Host", "Localhost:4000")
			return c.Next()
		}
		return c.Status(403).SendString("Request origin not allowed")
	})

	// Upgraded websocket request
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:4000"
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	// ws://localhost:4000/ws
	log.Fatal(app.Listen(":4000"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is backend web EL!")
	})

	api := app.Group("/api")
	route.CauHoiRoute(api.Group("/cau-hoi"))
	route.CauHoiKepRoute(api.Group("/cau-hoi-kep"))
	route.BothCauHoiRoute(api.Group("/both-cau-hoi"))
	route.UserRoute(api.Group("/users"))
	route.DangKienThucRoute(api.Group("/dkt"))
	route.DonViKienThucRoute(api.Group("/dvkt"))
	route.MoTaChiTietRoute(api.Group("/chi-tiet"))
	route.DanhMucRoute(api.Group("/danh-muc")) //API này đang lỗi
}
