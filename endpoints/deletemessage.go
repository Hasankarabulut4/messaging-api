package endpoints

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "messaging-api/employee" // Assuming this path is correct for your project
    "os"

    "github.com/gofiber/fiber/v2"
)

func DeleteMessage(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        err := employee.Error{
            Message: "Please enter id",
            Status:  400,
        }
        return c.Status(400).JSON(err)
    }
    log.Println("Entered id : " + id)

    var allDatas employee.Data 
    
    file, err := os.Open("database.json")
    if err != nil {
        log.Printf("Error opening database file: %v", err) 
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database file unavailable."})
    }
    defer file.Close()

    input, err := ioutil.ReadAll(file)
    if err != nil {
        log.Printf("Error reading file content: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error reading database content."})
    }
    log.Println("File content read successfully. Size:", len(input), "bytes")


    err = json.Unmarshal(input, &allDatas)
    if err != nil {
		err := employee.Error {
			Message: "Json parse error!",
			Status: 400,
		}
		return c.Status(400).JSON(err)
	}

    log.Printf("AllDatas unmarshaled: %+v", allDatas)

    var index int = -1

    for i, v := range allDatas.Messages {
        if v.ID == id {
            index = i
            break
        }
    }

    if index == -1 {
        err := employee.Error{
            Message: "Please enter a valid id",
            Status:  400,
        }
        return c.Status(400).JSON(err)
    }
    allDatas.Messages = append(allDatas.Messages[:index], allDatas.Messages[index+1:]...)

    jsonData, err := json.MarshalIndent(allDatas, "", "  ")
    if err != nil {
        log.Printf("Error marshaling updated data to JSON: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error preparing response data."})
    }

    err = os.WriteFile("database.json", jsonData, 0644)
    if err != nil {
        log.Printf("Error writing updated data to file: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving changes to database."})
    }
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Message deleted successfully"})
}