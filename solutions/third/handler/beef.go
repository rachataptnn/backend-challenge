package handler

import (
	"beefapi/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetBeefSummary(c *fiber.Ctx) error {
	beefSummary := service.GetBeefSummary()
	log.Println("Get Beef Summary Success")

	return c.JSON(fiber.Map{
		"message": "Get Beef Summary Success",
		"beef":    beefSummary,
	})
}

func Refresh(c *fiber.Ctx) error {
	err := service.ReadBeefFile("./src/beef.txt")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Refresh Success")

	return c.JSON(fiber.Map{
		"message": "Refresh Success",
	})
}
