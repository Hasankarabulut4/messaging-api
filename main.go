package main

import (
	"log"
	"messaging-api/endpoints"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();
	app.Get("/message", endpoints.GetAllMessages)

	app.Post("/message/send", endpoints.AddMessage)

	app.Delete("/message/delete/:id", endpoints.DeleteMessage)

	log.Fatal(app.Listen(":8000"))
}