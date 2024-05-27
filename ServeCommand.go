package proxy_shell

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ServeCommand(app *fiber.App, group fiber.Router, command Cmd) error {
	group.Get("/", func(ctx *fiber.Ctx) error {
		buffer, err := json.Marshal(command)
		if err != nil {
			return ctx.Status(503).JSON(buffer)
		}
		return ctx.Status(200).JSON(buffer)
	})
	log.Fatal(app.Listen(":21507"))
	return nil
}
