package main

import "github.com/gofiber/fiber/v2"

// GetBeefSummary returns the word counts as JSON
func GetBeefSummary(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"beef": beefCounts})
}
