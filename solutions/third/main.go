package main

import (
	"beefapi/handler"
	"beefapi/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	service.ReadBeefFile()

	// Start Fiber server
	app := fiber.New()

	beefAPI := app.Group("/beef")
	beefAPI.Patch("/refresh", handler.Refresh)
	beefAPI.Get("/summary", handler.GetBeefSummary)

	log.Println("Server is running on :3000")
	log.Fatal(app.Listen(":3000"))
}
