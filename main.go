package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	app.Get("/message", func (c *fiber.Ctx) error {

		return c.JSON(messages)
	})

	app.Post("/message/send", func (c *fiber.Ctx) error {
		var message DtoMessage
		if err := c.BodyParser(&message); err != nil {
			log.Fatal("Invalid body")
			errorMsg := Error{
				Status: 400, // Bad request
				Message: "Bad request",
			}
			return c.Status(400).JSON(errorMsg);
		}
		var msg Message = Message{
			ID: uuid.NewString(),
			Message: message.Message,
			Username: message.Username,
		}
		messages = append(messages, msg)
		return c.JSON(fiber.Map{
			"message":"Message successfully created!",
			"User":msg,
		})
	})

	log.Fatal(app.Listen(":8000"))
}
/*
{
"message":""
}
*/