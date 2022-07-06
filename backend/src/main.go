package main

import (
	"cli-tutor-backend/src/pkg/dockerclient"

	"github.com/gofiber/fiber/v2"
)

func main() {
	type Container struct {
		ID string
	}
	app := fiber.New()

	app.Get("/api/createcontainer", func(ctx *fiber.Ctx) error {
		type Container struct {
			ID string
		}
		containerid := dockerclient.CreateContainer()

		response := Container{
			ID: containerid,
		}
		return ctx.JSON(response)
	})

	app.Post("/api/startcontainer", func(ctx *fiber.Ctx) error {
		container := new(Container)
		err := ctx.BodyParser(container)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}
		dockerclient.RunContainer(container.ID)
		return ctx.SendStatus(200)
	})

	app.Post("/api/stopcontainer", func(ctx *fiber.Ctx) error {
		container := new(Container)
		err := ctx.BodyParser(container)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}
		dockerclient.StopContainer(container.ID)
		return ctx.SendStatus(200)
	})

	app.Listen(":8080")
}
