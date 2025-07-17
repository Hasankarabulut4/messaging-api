package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"messaging-api/employee"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetAllMessages(c *fiber.Ctx) error {
	var data employee.Data
	var messages []employee.Message

	file, err := os.Open("database.json")
	if err != nil {
		err := employee.Error{
            Message: "Database not found!",
            Status:  400,
        }
        return c.Status(400).JSON(err)
	}

	raw, err := ioutil.ReadAll(file);
	if err != nil {
		err := employee.Error{
            Message: "Database Permission denied!",
            Status:  400,
        }
        return c.Status(400).JSON(err)
	}

	err = json.Unmarshal(raw, &data);
	if err != nil {
		err := employee.Error{
            Message: "Database Permission denied!",
            Status:  400,
        }
        return c.Status(400).JSON(err)
	}

	messages = data.Messages;
	return c.Status(200).JSON(messages);
}