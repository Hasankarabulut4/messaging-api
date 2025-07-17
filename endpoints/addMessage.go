package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"messaging-api/employee"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddMessage(c *fiber.Ctx) error {

	var message employee.DtoMessage
	var datas employee.Data

	if err := c.BodyParser(&message); err != nil {
		err := employee.Error {
			Message: "Please add message to body",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}

	file, err := os.Open("database.json")
	if err != nil {
		err := employee.Error {
			Message: "Database not found!",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		err := employee.Error {
			Message: "File permission denied!",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}

	err = json.Unmarshal(raw, &datas)
	if err != nil {
		err := employee.Error {
			Message: "Json parse error!",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}
	msg := employee.Message{
		Message: message.Message,
		Username: message.Username,
		ID: strconv.Itoa(datas.NextId),
	}

	datas.Messages = append(datas.Messages, msg)
	datas.NextId = datas.NextId+1

	data, err := json.MarshalIndent(datas, "", " ");
	if err != nil {
		err := employee.Error {
			Message: "Json convert error!",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}

	os.WriteFile("database.json", data, 0644)

	return c.Status(200).JSON(fiber.Map{
		"inf":"Message succesfully sent!",
		"msg":msg,
	})
}