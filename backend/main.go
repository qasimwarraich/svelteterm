package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"cli-tutor-backend/pkg/dockerclient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	type Container struct {
		ID string
	}

	type ContainerResize struct {
		ID     string
		Height uint
		Width  uint
	}

	app := fiber.New(fiber.Config{BodyLimit: 100 * 1024 * 1024})
	app.Use(cors.New(cors.ConfigDefault))

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

	app.Post("/api/resizecontainer", func(ctx *fiber.Ctx) error {
		container := new(ContainerResize)
		err := ctx.BodyParser(container)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}
		dockerclient.HandlePty(container.ID, container.Height, container.Width)
		return ctx.SendStatus(200)
	})

	app.Post("/api/upload", func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		} else {
			ts := time.Now().UTC().Format(time.RFC3339Nano)
			err := ctx.SaveFile(file, fmt.Sprintf(os.Getenv("LOGPATH")+"%s-%s", ts, file.Filename))
			if err != nil {
				ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
				log.Println(err)
				return err
			} else {
				log.Println("A log file has been saved.")
			}
		}
		return ctx.SendStatus(200)
	})

	log.Fatal(app.ListenTLS(":"+os.Getenv("PORT"), os.Getenv("CERTPATH"), os.Getenv("CERTKEYPATH")))
}
