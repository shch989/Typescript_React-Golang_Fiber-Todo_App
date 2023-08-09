package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Done  bool   `json: "done"`
	Body  string `json: "body"`
}

func main() {
	fmt.Println("Hello world")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber!")
	})

	// todo 리스트 추가
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)
		if err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	// todo 체크 리스트
	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = !todos[i].Done
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":5000"))

}
