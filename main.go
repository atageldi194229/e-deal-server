package main

import (
	"log"

	"github.com/atageldi194229/e-deal-server/database"
	"github.com/atageldi194229/e-deal-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// db connect
	database.Connect()
	database.Automigrate()

	// fiber app init
	app := fiber.New()

	// setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// setup routes
	routes.Setup(app)

	// server listen
	log.Fatal(app.Listen(":3001"))
}
