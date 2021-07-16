package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pittalamadhuri/loginRadiusServer/handler"
	"log"
)

func main() {
	app := fiber.New()

	testGroup := app.Group("/radius")

	testGroup.Post("", handler.FindTop10)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err.Error())
	}

}
