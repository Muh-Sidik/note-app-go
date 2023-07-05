package main

import (
	"log"
	"note-app/src"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db := src.ConnectDB()

	r := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Note App",
	})

	r.Use(cors.New())

	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "homepage",
			"data":    "nil",
		})
	})

	r.Route("/note", src.RouteHandler(db), "note")

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "Not Found",
		})
	})

	r.Listen("0.0.0.0:5000")
}
