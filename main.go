package main

import (
	"log"
	"messaging-api/endpoints"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Message      string `json:"message"`
}

type DtoMessage struct {
	Username     string `json:"username"`
	Message      string `json:"message"`
}

type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

var (
	messages []Message;
)

func main() {
	app := fiber.New();
	app.Get("/message", endpoints.GetAllMessages)

	app.Post("/message/send", endpoints.AddMessage)

	app.Delete("/message/delete/:id", endpoints.DeleteMessage)

	log.Fatal(app.Listen(":8000"))
}