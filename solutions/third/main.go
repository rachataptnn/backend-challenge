package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Global map to store beef counts
var beefCounts map[string]int

func main() {
	// Load data from file
	beefCounts = LoadBeefData("./data.txt")

	// Start Fiber server
	app := fiber.New()
	app.Get("/beef/summary", GetBeefSummary)

	log.Println("Server is running on :3000")
	log.Fatal(app.Listen(":3000"))
}
